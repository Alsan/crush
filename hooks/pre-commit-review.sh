#!/usr/bin/env bash
set -euo pipefail

# Pre-commit code review hook for Crush.
# Blocks git commit commands and injects context instructing the agent
# to perform a code review before committing.

cmd="${CRUSH_TOOL_INPUT_COMMAND:-}"

# Only gate git commit commands (not git add, git status, etc.)
if ! echo "$cmd" | grep -qE '^\s*git\s+commit'; then
  exit 0
fi

# Allow commit --amend for quick fixes (no new code to review)
if echo "$cmd" | grep -qE '\-\-amend'; then
  exit 0
fi

# Extract changed files for the review context.
changed_files=$(git diff --cached --name-only 2>/dev/null || git diff --name-only 2>/dev/null || echo "")
if [ -z "$changed_files" ]; then
  exit 0
fi

cat <<'REVIEW_MSG'
Pre-commit code review required. Before committing:

1. Run `tokensave_diff_context` or `tokensave_commit_context` to understand what changed
2. Review the diff for: correctness, Go safety (nil/slice/concurrency), error handling, test coverage
3. Run `go build ./...` and `go test ./... -count=1` to verify
4. Only then retry the git commit

If this is a trivial/docs-only commit, use `git commit --amend` to bypass this check.
REVIEW_MSG

exit 2
