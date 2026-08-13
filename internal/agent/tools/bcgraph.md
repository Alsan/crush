# Bcgraph
Forward to a better_code_review_graph tool function and return its JSON result.

## Usage

```json
{
  "tool": "query_graph",
  "params": { "pattern": "callers_of", "target": "SomeFunction" },
  "repo_root": "/path/to/repo"
}
```

## Params

- `tool` (string, required): the better-code-review-graph function to call. One of:
  `build_or_update_graph`, `query_graph`, `review_delta`, `diff_graph`,
  `get_review_context`, `get_docs_section`, `get_impact_radius`,
  `list_graph_stats`, `semantic_search_nodes`, `find_large_functions`,
  `embed_graph`, `export_graph_dispatch`, `import_graph_dispatch`,
  `renamed_in_diff`.
- `repo_root` (string, optional): repository root; defaults to the working directory.
- `params` (object, optional): keyword arguments for the chosen tool.

Run-to-completion is bounded by `CRUSH_BCRG_TIMEOUT` (seconds, default 300); output is capped to keep memory bounded.
