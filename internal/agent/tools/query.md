# query

Run a predefined graph query against the better-code-review-graph knowledge graph.

## Usage

```json
{
  "pattern": "callers_of",
  "target": "SomeFunction",
  "repo_root": "/path/to/repo"
}
```

## Params

- `pattern` (string, required): `callers_of`, `callees_of`, `imports_of`, `importers_of`, `children_of`, `tests_for`, `inheritors_of`, `file_summary`.
- `target` (string, required): node name, qualified name, or file path to query.
- `repo_root` (string, optional): repository root; auto-detected if omitted.
