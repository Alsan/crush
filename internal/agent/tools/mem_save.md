# mem_save

Save an important observation to persistent memory across sessions.

## Usage

```json
{
  "title": "JWT auth middleware",
  "content": "**What**: ... \n**Why**: ... \n**Where**: ... \n**Learned**: ...",
  "type": "bugfix"
}
```

## Params

- `title` (string, required): short, searchable title.
- `content` (string, required): structured content (What/Why/Where/Learned).
- `type` (string, optional): decision, architecture, bugfix, pattern, config, discovery, learning, manual.
- `topic_key` (string, optional): stable key for upserts.
