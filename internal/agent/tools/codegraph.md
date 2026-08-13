# Codegraph
Native access to CodeGraph's CLI knowledge-graph index over this codebase.

## Usage

```json
{
  "op": "query",
  "query": "handleRequest",
  "path": "/path/to/repo",
  "limit": 10
}
```

## Params

- `op` (string, required): CodeGraph CLI subcommand:
  `query`, `callers`, `callees`, `context`, `files`, `impact`, `affected`,
  `index`, `sync`, `status`.
- `path` (string, optional): project path; defaults to the working directory.
- `query` / `symbol` (string, optional): the search term or symbol for
  `query`, `callers`, `callees`, `context`.
- `limit` (int, optional): max results.
- `files` (list, optional): changed files for `affected`.

Run-to-completion is bounded by `CRUSH_CODEGRAPH_TIMEOUT` (seconds, default 300); output is capped to keep memory bounded.
