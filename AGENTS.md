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
