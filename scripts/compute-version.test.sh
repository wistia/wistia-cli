#!/usr/bin/env bash
#
# Unit tests for compute-version.sh.
# Run: bash scripts/compute-version.test.sh
#
# Each case runs in a throwaway git repo with the given tags, invokes
# compute-version.sh with a branch name, and checks the output (or that it errors).
# Exits non-zero if any case fails.
set -uo pipefail

SCRIPT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/compute-version.sh"
pass=0
fail=0

# run_case <branch> <space-separated-tags> <expected>
# expected "ERROR" means the script should exit non-zero.
run_case() {
  local branch="$1" tags="$2" expect="$3" tmp out rc mark display
  tmp="$(mktemp -d)"
  out="$(
    cd "$tmp" || exit 99
    git init -q && git config user.email t@example.com && git config user.name t \
      && git commit -q --allow-empty -m init || exit 98
    for tg in $tags; do git tag "$tg" || exit 97; done
    "$SCRIPT" "$branch" 2>/dev/null
  )"
  rc=$?
  rm -rf "$tmp"

  if [ "$expect" = "ERROR" ]; then
    [ "$rc" -ne 0 ] && mark="PASS" || mark="FAIL"
    display="exit $rc"
  else
    { [ "$rc" -eq 0 ] && [ "$out" = "$expect" ]; } && mark="PASS" || mark="FAIL"
    display="$out"
  fi
  [ "$mark" = "PASS" ] && pass=$((pass + 1)) || fail=$((fail + 1))
  printf "%-4s branch=%-13s tags=[%-28s] => %-10s (want %s)\n" \
    "$mark" "$branch" "$tags" "$display" "$expect"
}

run_case "v2026-05"    ""                               "2026.5.0"   # first release
run_case "v2026-05"    "v2026.5.0 v2026.5.1 v2026.5.2"  "2026.5.3"   # increment highest patch
run_case "v2026-05"    "v2026.5.0 v2026.5.10 v2026.5.2" "2026.5.11"  # numeric, not lexical, max
run_case "v2026-05"    "v2026.04.8 v2026.04.5"          "2026.5.0"   # legacy padded tags ignored
run_case "v2026-01"    "v2026.1.0"                      "2026.1.1"   # different version branch
run_case "v2026-12"    ""                               "2026.12.0"  # double-digit month
run_case "main"        ""                               "ERROR"      # not a version branch
run_case "v2026-05-rc" ""                               "ERROR"      # malformed branch name

echo
echo "Results: ${pass} passed, ${fail} failed"
[ "$fail" -eq 0 ]
