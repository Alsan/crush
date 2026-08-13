# mem_context

Get recent memory context from previous sessions, scoped to this project.

## Usage

```json
{
  "scope": "project",
  "limit": 10
}
```

## Params

- `scope` (string, optional): "project" (default) or "personal".
- `limit` (int, optional): max recent memories (default unspecified; store decides).
