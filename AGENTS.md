# Crush Development Guide

Crush is a terminal-based AI coding assistant built in Go by
[Charm](https://charm.land). Module: `github.com/charmbracelet/crush`.

## Build / Test / Lint

| 命令 | 用途 |
|------|------|
| `task build` / `go build .` | 构建 |
| `task test` / `go test ./...` | 测试（单测：`go test ./pkg -run TestName`） |
| `task lint:fix` | Lint |
| `task fmt` | 格式化 (`gofumpt -w .`) |
| `task modernize` | 代码简化 |
| `task dev` | 开发模式（带 profiling） |
| `go test ./... -update` | 更新 golden files |

## On-Demand Reference

| 文档 | 内容 |
|------|------|
| `docs/development.md` | 架构、代码风格、测试模式、样式系统、提交规范 |
| `internal/ui/AGENTS.md` | TUI 开发指南 |
| `HOOKS.md` | 用户侧 hooks 协议 |
| `docs/hooks/README.md` | Hooks 技术文档 |

## Beads

全局规则见 `~/.config/crush/AGENTS.md` 与 `prompts/beads-workflow.md`
（命令速查、依赖方向、bv 分诊、会话关闭协议）——此处只写本项目差异：

- 后端：`bd` + Dolt，issue 前缀 `crush-*`
- 进入先 `bd ready`；完整上下文 `bd prime`
- 只读分诊可用 `bv --robot-triage`（勿裸跑，TUI 会卡住会话）

## Code Review Protocol (Pre-Commit)

`git commit` triggers a PreToolUse hook (`hooks/pre-commit-review.sh`) that
blocks the commit and injects review instructions. The hook:

1. Detects changed Go files in the staging area
2. Runs `go build`, `go vet`, `go test -short` to verify correctness
3. Classifies changes and recommends relevant golang-* skills
4. Blocks commit (exit 2) with a structured review checklist

**When the hook fires**, follow this workflow:

1. **Load recommended skills** — the hook lists which golang-* skills to load
   via `mcp_superpowers_use_skill`. Always start with `golang-how-to` as the
   orchestrator; it routes to the right secondary skills.
2. **Get context** — run `tokensave_commit_context` or `tokensave_diff_context`
   to understand what changed.
3. **Review against the checklist** — correctness, concurrency, style, testing,
   security, performance (see hook output for details).
4. **Fix issues or bypass**:
   - Issues found → fix, then retry `git commit`
   - All clear → `git commit --amend` (bypasses the hook)
   - Trivial/docs-only → `git commit --amend --allow-empty`

**Bypass paths** (no review needed):
- `git commit --amend` — existing commit, no new code
- `git commit --allow-empty` — merge commits
- Non-Go file changes only

**Skill loading reference** (from `golang-how-to`):

| Change type | Skills to load |
|-------------|---------------|
| Error handling | `golang-error-handling` + `golang-safety` |
| Concurrency | `golang-concurrency` + `golang-context` |
| CLI commands | `golang-cli` + `golang-spf13-cobra` |
| Database | `golang-database` + `golang-error-handling` |
| Security | `golang-security` + `golang-safety` |
| Performance | `golang-performance` + `golang-benchmark` |
| Tests | `golang-testing` + `golang-stretchr-testify` |
| Style/naming | `golang-code-style` + `golang-naming` + `golang-lint` |
