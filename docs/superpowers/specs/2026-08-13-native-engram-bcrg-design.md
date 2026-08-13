# Design: Native engram + better-code-review-graph integration

Date: 2026-08-13
Status: Approved (personal-only build)
Scope: Expose engram memory tools and better-code-review-graph tools as native
Crush builtin tools, not via MCP.

## Context

Crush currently reaches engram (`mem_*`) and better-code-review-graph
(`graph`/`query`/`review`) only through MCP servers configured in
`crush.json`. The user wants these as native Crush builtin tools.

Both upstream tools are MCP-only in their own repos: the mem_* operations run
in a separate MCP (stdio) process, and bcrg's tool logic lives in Python.
Native integration must therefore do one of:

- engram: embed its Go store in-process (user chose this — "embed_go")
- bcrg: exec the Python CLI and parse JSON output (user chose this — "exec_python")

This build is **personal-only**, so a local Go workspace (`go.work`) is
acceptable (user chose option A).

## Goals

- Add Crush builtin tools: `mem_save`, `mem_search`, `mem_context` (engram),
  and `graph`, `query`, `review` (better-code-review-graph).
- Engram tools run in-process against engram's store (no MCP, no subprocess).
- Bcrg tools exec the `better-code-review-graph` Python CLI and parse its
  JSON output.
- All tools follow Crush's existing tool pattern (`bash.go`: a `Params`
  struct + `New*Tool(...)` constructor + a `.md` description), go through the
  permission service, and are registered in `coordinator.go:buildTools`.

## Non-goals

- Not exposing every engram `mem_*` tool (only the 3 core ones initially).
- Not reimplementing bcrg's graph engine in Go.
- Not making this a shared/CI build (personal-only; no release concern).

## Architecture

### 1. Go workspace wiring

Create `go.work` at the Crush repo root that includes both modules:

```go
go 1.25

use (
  .
  /Users/mac/data/oss/engram
)
```

Crush's module is `github.com/charmbracelet/crush`; engram's is
`github.com/Gentleman-Programming/engram`.

### 2. engram public `mem` package

Engram's store lives in `internal/store` (not importable across modules). Add
a public package to the engram repo, `github.com/Gentleman-Programming/engram/mem`,
that wraps the store and exposes a small, Crush-consumable API:

- `mem.Open(cfg)` → returns a handle to a store for a given data dir.
- `mem.Save(...)` — write an observation (mirrors `mem_save`).
- `mem.Search(...)` — query observations (mirrors `mem_search`).
- `mem.Context(...)` — recent context (mirrors `mem_context`).

This package reuses `internal/store` internally. Signature details to be
finalized in the implementation plan against engram's actual store API
(`store.New(Config{DataDir, ...})`).

### 3. Crush builtin tools

Under `internal/agent/tools/`:

- `mem_save.go` + `mem_save.md`
- `mem_search.go` + `mem_search.md`
- `mem_context.go` + `mem_context.md`
- `graph.go` + `graph.md`
- `query.go` + `query.md`
- `review.go` + `review.md`

Each engram tool imports the engram `mem` package and calls it in-process.
Each bcrg tool execs `better-code-review-graph <action> ...` and parses the
JSON on stdout.

### 4. Registration and permissions

Register all six in `coordinator.go` `buildTools(...)` (alongside the existing
`NewBashTool` etc.). If an engram tool touches a path (e.g. a project data
dir), it passes that to `permission.Service.Request`; otherwise it relies on
the tool-level allowlist.

### 5. Data directory

Engram store needs a data dir. Use Crush's existing per-project data directory
(`options.data_directory`, default `.crush`), e.g.
`<data_directory>/engram`. Ensures isolation per project and reuses Crush's
config plumbing.

## Component breakdown

| Component | Responsibility | Depends on |
|-----------|----------------|-----------|
| `engram/mem` (upstream) | Public API over engram store | `engram` internal/store |
| `tools/mem_*.go` | Crush tool adapters (in-process) | `engram/mem`, `permission` |
| `tools/graph.go` etc. | Crush tool adapters (exec Python CLI) | `os/exec`, `permission` |
| `go.work` | Build wiring | both modules |

## Data flow

`user/agent -> tool call -> mem_save.go -> engram/mem.Save -> store`
`user/agent -> tool call -> graph.go -> exec better-code-review-graph graph -> JSON parsing -> result`

## Error handling

- engram tools: propagate store errors as tool errors.
- bcrg tools: if the binary is missing or exits non-zero, surface the stderr
  as a readable error. Wrap CLI `--help`/version probes lazily.

## Testing

- Add unit tests for each builtin tool under
  `internal/agent/tools/` following existing `*_test.go` patterns
  (bcrg wrappers can use a mock/fake binary; engram can use a temp data dir).
- `go build ./...` and `task test` must pass with `go.work` in place.

## Open items (resolved)

- Backend: option A (local go.work). Confirmed.
- Tool list: the 6 listed core tools. Confirmed.
- bcrg integration: exec Python CLI (not Go reimplementation). Confirmed.
