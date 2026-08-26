#!/usr/bin/env bash
set -euo pipefail

# Pre-commit code review hook for Crush (Go project).
# Blocks git commit and injects golang-* skill review instructions.

cmd="${CRUSH_TOOL_INPUT_COMMAND:-}"

# Only gate git commit commands (not git add, git status, etc.)
if ! echo "$cmd" | grep -qE '^\s*git\s+commit'; then
  exit 0
fi

# Allow commit --amend for quick fixes (no new code to review)
if echo "$cmd" | grep -qE '\-\-amend'; then
  exit 0
fi

# Allow empty commits (e.g. merge commits with no diff)
if echo "$cmd" | grep -qE '\-\-allow-empty'; then
  exit 0
fi

# Extract changed files for the review context.
changed_files=$(git diff --cached --name-only 2>/dev/null || git diff --name-only 2>/dev/null || echo "")
if [ -z "$changed_files" ]; then
  exit 0
fi

# Check if any Go files are in the diff.
go_files=$(echo "$changed_files" | grep -E '\.go$' || true)
if [ -z "$go_files" ]; then
  # Non-Go changes: allow commit without Go review
  exit 0
fi

# Classify changes to determine relevant golang-* skills.
has_concurrency=$(echo "$go_files" | xargs grep -l 'goroutine\|chan \|sync\.\|context\.' 2>/dev/null || true)
has_cli=$(echo "$go_files" | xargs grep -l 'cobra\|urfave' 2>/dev/null || true)
has_grpc=$(echo "$go_files" | xargs grep -l 'grpc\|proto' 2>/dev/null || true)
has_db=$(echo "$go_files" | xargs grep -l 'sql\.\|database\|pgx\|sqlx' 2>/dev/null || true)
has_security=$(echo "$go_files" | xargs grep -l 'crypto\|tls\|jwt\|token\|secret\|password' 2>/dev/null || true)
has_perf=$(echo "$go_files" | xargs grep -l 'pprof\|benchmark\|sync.Pool\|make(chan' 2>/dev/null || true)
has_test=$(echo "$go_files" | grep -E '_test\.go$' || true)

# Count changed Go files for context.
go_file_count=$(echo "$go_files" | wc -l | tr -d ' ')

# Build the changed files list for context.
file_list=$(echo "$go_files" | head -20)

# --- Build skill recommendation list ---
skills="golang-how-to, golang-code-style, golang-safety, golang-error-handling, golang-testing"

if [ -n "$has_concurrency" ]; then
  skills="$skills, golang-concurrency, golang-context"
fi
if [ -n "$has_cli" ]; then
  skills="$skills, golang-cli, golang-spf13-cobra"
fi
if [ -n "$has_grpc" ]; then
  skills="$skills, golang-grpc"
fi
if [ -n "$has_db" ]; then
  skills="$skills, golang-database"
fi
if [ -n "$has_security" ]; then
  skills="$skills, golang-security"
fi
if [ -n "$has_perf" ]; then
  skills="$skills, golang-performance, golang-benchmark"
fi
if [ -n "$has_test" ]; then
  skills="$skills, golang-stretchr-testify"
fi

# --- Run build and test verification ---
build_ok=true
test_ok=true
vet_ok=true

if ! go build ./... 2>/dev/null; then
  build_ok=false
fi

if ! go vet ./... 2>/dev/null; then
  vet_ok=false
fi

if ! go test ./... -count=1 -short 2>/dev/null; then
  test_ok=false
fi

# --- Build the review message ---
cat >&2 <<REVIEW_MSG
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  PRE-COMMIT CODE REVIEW REQUIRED (Go project)
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Changed Go files (${go_file_count}):
${file_list}

Verification:
  go build ./...       : $([ "$build_ok" = true ] && echo "✅ PASS" || echo "❌ FAIL")
  go vet ./...         : $([ "$vet_ok" = true ] && echo "✅ PASS" || echo "❌ FAIL")
  go test -short ./... : $([ "$test_ok" = true ] && echo "✅ PASS" || echo "❌ FAIL")

━━━ REQUIRED: Load these golang-* skills BEFORE reviewing ━━━

Load skills (via mcp_superpowers_use_skill or by reading their SKILL.md):
  ${skills}

━━━ Review Checklist ━━━

1. CONTEXT: Run tokensave_commit_context or tokensave_diff_context to
   understand what changed and why.

2. CORRECTNESS:
   - Error handling: every error returned, wrapped with %w, not swallowed
   - Nil safety: pointers checked before dereference
   - Slice/map: preallocate capacity, check nil before append
   - Resource lifecycle: defer close, no goroutine leaks

3. CONCURRENCY (if applicable):
   - No data races: shared state protected by mutex or channel
   - Context propagation: passed as first param, cancellation respected
   - Goroutine lifecycle: bounded, cancellable, not leaked

4. STYLE (golang-code-style, golang-naming):
   - goimports grouping: stdlib / external / internal
   - Naming: PascalCase exported, camelCase unexported
   - Comments end with periods, wrap at 78 cols
   - Log messages start with capital letter

5. TESTING (golang-testing):
   - New code has corresponding tests
   - Table-driven tests with t.Parallel()
   - Use testify require (not assert) for critical checks
   - Use t.Tempdir() for temp directories

6. SECURITY (if applicable):
   - No hardcoded secrets/tokens
   - Input validation on external data
   - SQL injection prevention (parameterized queries)
   - TLS/crypto usage is correct

7. PERFORMANCE (if applicable):
   - No unnecessary allocations in hot paths
   - sync.Pool for frequently allocated objects
   - String concatenation uses strings.Builder

━━━ Workflow ━━━

After loading skills and reviewing:
  a) If issues found → fix them, then retry commit
  b) If all clear → use git commit --amend to bypass this check
     (the hook allows amend to skip re-review)

For trivial/docs-only Go changes (comments, imports only), use:
  git commit --amend --allow-empty

REVIEW_MSG

exit 2
