#!/usr/bin/env bash
#
# Smoke-test the staged npm packages via hermetic local installs.
#
# Usage: scripts/npm-smoke.sh --version <2026.M.N> --tarballs <dir>
#
# Every install runs --offline with a fresh cache so the registry is never
# consulted (the @wistia/* packages may not exist there yet). Three cases:
#   1. meta + matching platform tarball -> `wistia version` works
#   2. same, with --ignore-scripts      -> identical (there are no scripts)
#   3. meta only, --omit=optional       -> shim fails with actionable guidance
set -euo pipefail

VERSION=""
TARBALLS=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    --version) VERSION="$2"; shift 2 ;;
    --tarballs) TARBALLS="$2"; shift 2 ;;
    *) echo "usage: $0 --version <v> --tarballs <dir>" >&2; exit 1 ;;
  esac
done
[[ -n "$VERSION" && -n "$TARBALLS" ]] || { echo "usage: $0 --version <v> --tarballs <dir>" >&2; exit 1; }
TARBALLS=$(cd "$TARBALLS" && pwd)

key=$(node -p 'process.platform + "-" + process.arch')
meta_tgz="${TARBALLS}/wistia-wistia-cli-${VERSION}.tgz"
platform_tgz="${TARBALLS}/wistia-wistia-cli-${key}-${VERSION}.tgz"
[[ -f "$meta_tgz" ]] || { echo "error: ${meta_tgz} not found" >&2; exit 1; }
[[ -f "$platform_tgz" ]] || { echo "error: no platform tarball for ${key}" >&2; exit 1; }

hermetic_install() {
  local workdir="$1"
  shift
  (cd "$workdir" \
    && npm init -y >/dev/null 2>&1 \
    && npm install "$@" --offline --cache "$(mktemp -d)" \
        --no-audit --no-fund --loglevel=error >/dev/null)
}

echo "case 1: install meta + platform (${key}), run the binary"
t1=$(mktemp -d)
hermetic_install "$t1" "$meta_tgz" "$platform_tgz"
out=$("${t1}/node_modules/.bin/wistia" version)
grep -F "$VERSION" <<<"$out" >/dev/null || { echo "FAIL: 'wistia version' printed '${out}', expected ${VERSION}" >&2; exit 1; }
echo "PASS: ${out}"

echo "case 2: same install with --ignore-scripts"
t2=$(mktemp -d)
hermetic_install "$t2" "$meta_tgz" "$platform_tgz" --ignore-scripts
out=$("${t2}/node_modules/.bin/wistia" version)
grep -F "$VERSION" <<<"$out" >/dev/null || { echo "FAIL: 'wistia version' printed '${out}', expected ${VERSION}" >&2; exit 1; }
echo "PASS: ${out}"

echo "case 3: meta only with --omit=optional, expect an actionable failure"
t3=$(mktemp -d)
hermetic_install "$t3" "$meta_tgz" --omit=optional
if [[ -d "${t3}/node_modules/@wistia/wistia-cli-${key}" ]]; then
  echo "FAIL: platform package installed despite --omit=optional" >&2
  exit 1
fi
set +e
out=$("${t3}/node_modules/.bin/wistia" version 2>&1)
rc=$?
set -e
[[ $rc -ne 0 ]] || { echo "FAIL: expected nonzero exit without the platform package" >&2; exit 1; }
grep -qi "optional" <<<"$out" || { echo "FAIL: error lacks optional-deps guidance: ${out}" >&2; exit 1; }
grep -q "cli-installation" <<<"$out" || { echo "FAIL: error lacks the install-docs link: ${out}" >&2; exit 1; }
echo "PASS: exit ${rc}, guidance present"

echo "smoke: all cases passed"
