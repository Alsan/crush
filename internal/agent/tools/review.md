# review

Review what changed between two commit SHAs using the better-code-review-graph.

## Usage

```json
{
  "repo_root": "/path/to/repo",
  "from_sha": "abc123",
  "to_sha": "def456"
}
```

## Params

- `repo_root` (string, optional): repository root; auto-detected if omitted.
- `from_sha` (string, optional): starting commit SHA (default empty).
- `to_sha` (string, optional): ending commit SHA.
