# mem_search

Search past memories/persistent knowledge across sessions.

## Usage

```json
{
  "query": "JWT auth middleware",
  "type": "decision",
  "limit": 10
}
```

## Params

- `query` (string, required): search text, natural language or keywords.
- `type` (string, optional): filter by observation type (decision, architecture, bugfix, pattern, config, discovery, learning, manual).
- `match_mode` (string, optional): token matching "all" (default) or "any".
- `limit` (int, optional): max results (default 10).
