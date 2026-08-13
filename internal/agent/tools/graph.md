# graph

Build or update the better-code-review-graph code knowledge graph for a repository.

Run-to-completion is bounded by `CRUSH_BCRG_TIMEOUT` (seconds, default 300);
a timed-out build is killed and reported as an error.

## Usage

```json
{
  "repo_root": "/path/to/repo",
  "full_rebuild": false,
  "base": "HEAD~1"
}
```

## Params

- `repo_root` (string, optional): repository root; auto-detected if omitted.
- `full_rebuild` (bool, optional): re-parse every file (default false = incremental).
- `base` (string, optional): git ref for incremental diff (default `HEAD~1`).
