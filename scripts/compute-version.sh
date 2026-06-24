#!/usr/bin/env bash
#
# Compute the next dated release version for a version branch.
#
# Usage: scripts/compute-version.sh [branch-name]
#   branch-name defaults to $GITHUB_REF_NAME.
#
# A version branch is named v<year>-<month> (zero-padded month), e.g. v2026-05.
# The release version is unpadded dated SemVer <year>.<month>.<patch>, e.g.
# 2026.5.N, where N is the highest existing v<year>.<month>.* tag patch + 1
# (0 when no such tag exists). The dated version lives only on the git tag /
# GoReleaser ldflags, never in the Go module (see F10 in the release plan).
#
# Requires tags to be fetched first (git fetch --tags).
set -euo pipefail

branch="${1:-${GITHUB_REF_NAME:-}}"

if [[ ! "$branch" =~ ^v(20[0-9]{2})-([0-9]{2})$ ]]; then
  echo "error: expected a version branch like v2026-05, got '${branch}'" >&2
  exit 1
fi

year="${BASH_REMATCH[1]}"
month=$((10#${BASH_REMATCH[2]}))   # strip leading zero, force base-10
base="${year}.${month}"

# Highest patch among unpadded v<base>.* tags (e.g. v2026.5.3 -> 3).
# Legacy padded tags (e.g. v2026.04.8) do not match v<base>.* and are ignored.
highest=-1
while IFS= read -r tag; do
  [[ "$tag" =~ ^v${year}\.${month}\.([0-9]+)$ ]] || continue
  patch="${BASH_REMATCH[1]}"
  (( patch > highest )) && highest=$patch
done < <(git tag -l "v${base}.*")

echo "${base}.$(( highest + 1 ))"
