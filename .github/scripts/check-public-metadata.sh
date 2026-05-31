#!/usr/bin/env sh
set -eu

blocked_name_pattern='(^|[[:space:]_./-])(codex|claude|gemini|copilot|cursor|windsurf|aider|devin)([[:space:]_./-]|$)'
failed=0

check_value() {
  label="$1"
  value="${2:-}"

  if [ -z "$value" ]; then
    return 0
  fi

  if printf '%s\n' "$value" | grep -Eiq "$blocked_name_pattern"; then
    echo "::error title=Blocked coding tool name::${label} contains a blocked coding tool name. Use a product, issue, or behavior-focused name instead."
    failed=1
  fi
}

collect_commit_messages() {
  output_file="$1"
  : > "$output_file"

  if [ -n "${PRMAVEN_COMMIT_MESSAGES_FILE:-}" ]; then
    cat "$PRMAVEN_COMMIT_MESSAGES_FILE" > "$output_file"
    return 0
  fi

  if ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    return 0
  fi

  base_sha="${PRMAVEN_BASE_SHA:-}"
  head_sha="${PRMAVEN_HEAD_SHA:-}"
  push_before="${PRMAVEN_PUSH_BEFORE:-}"
  push_after="${PRMAVEN_PUSH_AFTER:-}"

  if [ -n "$base_sha" ] && [ -n "$head_sha" ] &&
    git cat-file -e "$base_sha^{commit}" 2>/dev/null &&
    git cat-file -e "$head_sha^{commit}" 2>/dev/null; then
    git log --format=%B --no-merges "$base_sha..$head_sha" > "$output_file" || :
    return 0
  fi

  if [ -n "$push_before" ] &&
    [ "$push_before" != "0000000000000000000000000000000000000000" ] &&
    [ -n "$push_after" ] &&
    git cat-file -e "$push_before^{commit}" 2>/dev/null &&
    git cat-file -e "$push_after^{commit}" 2>/dev/null; then
    git log --format=%B --no-merges "$push_before..$push_after" > "$output_file" || :
    return 0
  fi

  git log --format=%B --no-merges -n 50 > "$output_file" || :
}

check_value "Pull request title" "${PRMAVEN_PR_TITLE:-}"
check_value "Branch name" "${PRMAVEN_HEAD_REF:-}"

messages_file="$(mktemp)"
collect_commit_messages "$messages_file"

if [ -s "$messages_file" ] && grep -Eiq "$blocked_name_pattern" "$messages_file"; then
  echo "::error title=Blocked coding tool name::Commit messages contain a blocked coding tool name. Use product, issue, or behavior-focused commit messages instead."
  failed=1
fi

rm -f "$messages_file"

if [ "$failed" -ne 0 ]; then
  cat <<'MSG'
Public repository metadata must not include coding agent or tool names.

Rename the branch, pull request title, or commit message so the project history
describes the product change rather than the automation used to produce it.
MSG
  exit 1
fi

echo "Public metadata naming guard passed."
