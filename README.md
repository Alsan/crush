# Crush (Personal Fork)

> Based on [charmbracelet/crush](https://github.com/charmbracelet/crush).
> Compiles from source only.

## Tools Integrated

| Tool | Purpose |
|------|---------|
| **Pre-commit Code Review** | Blocks `git commit` on Go changes, loads golang-* skills for review |
| **git-cliff** | Auto-generates CHANGELOG.md on push |
| **Beads (bd)** | Issue tracking (replaces markdown TODOs) |
| **Codegraph / tokensave** | Code intelligence over indexed knowledge graph |
| **RTK** | Token compression for CLI output |
| **bcgraph** | Code review knowledge graph |

## Features Added

- **ctrl+/ editor keys help dialog** — quick reference for editor keybindings
- **Pre-commit golang-* skill review** — hook detects change patterns (concurrency, CLI, gRPC, DB, security, perf) and recommends matching skills
- **Strengthened file-read/search tool routing** — `view` as last-resort reader, clearer tool priority
- **Tightened coder system prompt** — ~30% smaller while preserving all rules
- **RTK CLI output routing** — compresses token usage
- **Beads integration** — `bd ready` / `bd claim` workflow for task tracking

## Fixes

- Keyed literals for `HelpRow` to satisfy `go vet`
- Remove vim-style single-letter keybindings from chat view
- Change select-all keybinding from `ctrl+shift+a` to `ctrl+a`
- Clarify file reads use `view`, not `read`; skip VCR tests without API key
- Clean rebase conflict leftovers in `go.sum` and editor keys
- Repair MCP stdio startup diagnostics
- Fix wrapped messages counted as newlines

## Build

```bash
go build .
# or
task build
```

## License

[FSL-1.1-MIT](https://github.com/charmbracelet/crush/raw/main/LICENSE.md)
