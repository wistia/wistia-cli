#!/usr/bin/env bash
#
# Publish the staged npm packages (see scripts/npm-stage.mjs).
#
# Env:
#   VERSION   (required)          release version, e.g. 2026.5.1
#   PKG_DIR   (default dist/npm)  staging dir containing manifest.json
#   CHANNEL   (default derived)   dist-tag channel; api-<year>-<padded month>
#   DRY_RUN   (default true)      when true, npm publish --dry-run, no dist-tags
#
# Publishes in manifest order (platform packages first, meta last) so the
# meta package never references versions that don't exist. Idempotent:
# already-published versions are skipped, so a partial failure can be re-run.
# Auth is ambient — a CI token via NODE_AUTH_TOKEN, or an interactive npm
# login session (expect an OTP prompt per publish).
set -euo pipefail

VERSION="${VERSION:?VERSION is required (e.g. 2026.5.1)}"
PKG_DIR="${PKG_DIR:-dist/npm}"
DRY_RUN="${DRY_RUN:-true}"
META_PKG="@wistia/wistia-cli"

if [[ ! "$VERSION" =~ ^[0-9]{4}\.[0-9]{1,2}\.[0-9]+$ ]]; then
  echo "error: VERSION must look like 2026.5.1, got '${VERSION}'" >&2
  exit 1
fi

if [[ -z "${CHANNEL:-}" ]]; then
  year="${VERSION%%.*}"
  rest="${VERSION#*.}"
  CHANNEL=$(printf 'api-%s-%02d' "$year" "${rest%%.*}")
fi

MANIFEST="${PKG_DIR}/manifest.json"
if [[ ! -f "$MANIFEST" ]]; then
  echo "error: ${MANIFEST} not found (run scripts/npm-stage.mjs first)" >&2
  exit 1
fi

mf() {
  node -p "require(require('path').resolve('${MANIFEST}'))$1"
}

# 0 = published, 1 = not published; aborts if the registry can't be reached.
already_published() {
  local name="$1" out attempt
  for attempt in 1 2 3; do
    if out=$(npm view "${name}@${VERSION}" version 2>&1); then
      return 0
    fi
    if grep -qE "E404|404 Not Found" <<<"$out"; then
      return 1
    fi
    echo "npm view ${name}@${VERSION} failed (attempt ${attempt}), retrying" >&2
    sleep $((attempt * 5))
  done
  echo "error: could not determine publish state for ${name}@${VERSION}: ${out}" >&2
  exit 1
}

count=$(mf .length)
for ((i = 0; i < count; i++)); do
  name=$(mf "[$i].name")
  dir=$(mf "[$i].dir")
  staged=$(mf "[$i].version")
  if [[ "$staged" != "$VERSION" ]]; then
    echo "error: manifest entry ${name} is staged at ${staged}, expected ${VERSION}" >&2
    exit 1
  fi
  if already_published "$name"; then
    echo "skip: ${name}@${VERSION} already published"
    continue
  fi
  publish_args=(publish "$dir" --access public --tag "$CHANNEL")
  [[ "$DRY_RUN" == "true" ]] && publish_args+=(--dry-run)
  echo "+ npm ${publish_args[*]}"
  npm "${publish_args[@]}"
done

if [[ "$DRY_RUN" == "true" ]]; then
  echo "dry run: skipping the latest dist-tag guard"
  exit 0
fi

# Move `latest` only when this version is the highest stable published —
# a backpatch to an older API version must not steal `latest` from a newer
# one. An empty result means first publish: this version is the max.
highest=$(npm view "$META_PKG" versions --json 2>/dev/null | node -p '
  const input = require("fs").readFileSync(0, "utf8").trim();
  const parsed = input ? JSON.parse(input) : [];
  const versions = Array.isArray(parsed) ? parsed : [parsed];
  versions
    .filter((v) => !v.includes("-"))
    .sort((a, b) => {
      const [x, y] = [a, b].map((v) => v.split(".").map(Number));
      return x[0] - y[0] || x[1] - y[1] || x[2] - y[2];
    })
    .pop() ?? ""')

if [[ -z "$highest" || "$highest" == "$VERSION" ]]; then
  for ((i = 0; i < count; i++)); do
    name=$(mf "[$i].name")
    npm dist-tag add "${name}@${VERSION}" latest
  done
  echo "latest -> ${VERSION} (all packages)"
else
  echo "latest stays at ${highest}; ${VERSION} published on ${CHANNEL} only"
fi
