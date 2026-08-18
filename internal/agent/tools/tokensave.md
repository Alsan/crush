Query the tokensave code-graph database (`.tokensave/tokensave.db`) for
this repository. The graph is prebuilt by the tokensave MCP server and
contains every indexed symbol (functions, methods, structs, traits,
classes, enums) plus their call/uses/implements relationships.

Prefer this tool over reading source files or running grep when you need
to locate symbols, trace call chains, or assess the impact of a change —
it answers structural questions in one query.

- `op` (string, required): one of
  - `status`: graph stats (node, edge, file counts, DB size)
  - `files`: indexed files with node counts (`file` filters by path prefix)
  - `entities`: table of contents for one file (`file` required)
  - `search`: full-text symbol search ranked by relevance (`query` required)
  - `find`: exact-match symbol name lookup (`name` required)
  - `node`: full node details by `id`
  - `body`: source code of a symbol, by `id` or `name`
  - `callers`: functions that call this node (`id`, optional `max_depth`
    for transitive callers)
  - `callees`: functions this node calls (`id`, optional `max_depth`)
  - `impact`: everything that transitively depends on this node (`id`)
- `root` (string, optional): project root containing `.tokensave/`;
  defaults to the working directory
- `id` (string, optional): node ID from a previous search/find result
- `name` (string, optional): exact symbol name (for `find`/`body`)
- `query` (string, optional): search terms (for `search`)
- `file` (string, optional): project-relative file path (for
  `entities`/`files`)
- `kinds` (string, optional): comma-separated node kinds filter
  (function, method, struct, enum, trait, class, interface, field, const)
- `edge_kinds` (string, optional): comma-separated edge kinds for graph
  ops; defaults to `calls` (others: uses, implements, extends, contains)
- `max_depth` (int, optional): traversal depth for `callers`/`callees`/
  `impact` (default 1, 1, 3; max 10)
- `limit` (int, optional): max results (default 20, max 200)

Typical flow: `search` or `find` to get a node ID, then `callers`,
`callees`, `impact`, or `body` with that ID. If the DB is missing, tell
the user to initialize the tokensave MCP server for this repository.
