package tools

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

// tokensaveFixture builds a temporary project with a minimal tokensave
// graph database and returns its root.
func tokensaveFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	dbDir := filepath.Join(root, ".tokensave")
	require.NoError(t, os.MkdirAll(dbDir, 0o755))
	dbPath := filepath.Join(dbDir, "tokensave.db")

	db, err := sql.Open("sqlite", "file:"+dbPath)
	require.NoError(t, err)
	defer db.Close()

	schema := `
CREATE TABLE nodes (
	id TEXT PRIMARY KEY,
	kind TEXT NOT NULL,
	name TEXT NOT NULL,
	qualified_name TEXT NOT NULL,
	file_path TEXT NOT NULL,
	start_line INTEGER NOT NULL,
	end_line INTEGER NOT NULL,
	docstring TEXT,
	signature TEXT,
	visibility TEXT NOT NULL DEFAULT 'private',
	is_async INTEGER NOT NULL DEFAULT 0,
	cognitive_complexity INTEGER NOT NULL DEFAULT 0
);
CREATE TABLE edges (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	source TEXT NOT NULL,
	target TEXT NOT NULL,
	kind TEXT NOT NULL,
	line INTEGER
);
CREATE TABLE files (
	path TEXT PRIMARY KEY,
	content_hash TEXT NOT NULL,
	size INTEGER NOT NULL,
	modified_at INTEGER NOT NULL,
	indexed_at INTEGER NOT NULL,
	node_count INTEGER NOT NULL DEFAULT 0,
	kind TEXT NOT NULL DEFAULT 'code'
);
CREATE VIRTUAL TABLE nodes_fts USING fts5(
	name, qualified_name, docstring, signature, search_terms,
	content='nodes', content_rowid='rowid',
	tokenize='porter unicode61'
);`
	_, err = db.Exec(schema)
	require.NoError(t, err)

	_, err = db.Exec(`INSERT INTO nodes (id, kind, name, qualified_name, file_path, start_line, end_line, signature)
VALUES
	('n-foo', 'function', 'Foo', 'pkg.Foo', 'a.go', 3, 5, 'func Foo()'),
	('n-bar', 'function', 'Bar', 'pkg.Bar', 'a.go', 7, 8, 'func Bar()'),
	('n-baz', 'struct', 'Baz', 'pkg.Baz', 'b.go', 1, 2, 'type Baz struct');`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO nodes_fts (rowid, name, qualified_name, docstring, signature, search_terms)
	SELECT rowid, name, qualified_name, COALESCE(docstring,''), COALESCE(signature,''), '' FROM nodes`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO edges (source, target, kind, line) VALUES ('n-foo', 'n-bar', 'calls', 4)`)
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO files (path, content_hash, size, modified_at, indexed_at, node_count)
VALUES ('a.go', 'h1', 10, 0, 0, 2), ('b.go', 'h2', 10, 0, 0, 1)`)
	require.NoError(t, err)

	source := "package main\n\nfunc Foo() {\n\tBar()\n}\n\nfunc Bar() {\n}\n"
	require.NoError(t, os.WriteFile(filepath.Join(root, "a.go"), []byte(source), 0o644))
	return root
}

func TestTokensaveSearchAndFind(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)
	ctx := context.Background()

	out, err := runTokensave(ctx, root, &TokensaveParams{Op: "search", Query: "Foo"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-foo"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "search", Query: "fnord-missing"})
	require.NoError(t, err)
	require.Contains(t, out, "null")

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "find", Name: "Bar"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-bar"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "find", Name: "Baz", Kinds: "struct"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-baz"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "find", Name: "Baz", Kinds: "function"})
	require.NoError(t, err)
	require.Contains(t, out, "null")
}

func TestTokensaveGraphTraversal(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)
	ctx := context.Background()

	out, err := runTokensave(ctx, root, &TokensaveParams{Op: "callers", ID: "n-bar"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-foo"`)
	require.Contains(t, out, `"depth": 1`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "callees", ID: "n-foo"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-bar"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "impact", ID: "n-bar"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-foo"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "callers", ID: "n-foo"})
	require.NoError(t, err)
	require.Contains(t, out, "null")
}

func TestTokensaveNodeAndBody(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)
	ctx := context.Background()

	out, err := runTokensave(ctx, root, &TokensaveParams{Op: "node", ID: "n-foo"})
	require.NoError(t, err)
	require.Contains(t, out, `"qualified_name": "pkg.Foo"`)
	require.Contains(t, out, `"incoming_edges": 0`)
	require.Contains(t, out, `"outgoing_edges": 1`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "body", ID: "n-foo"})
	require.NoError(t, err)
	require.Contains(t, out, `"from_line": 3`)
	require.Contains(t, out, "func Foo() {")

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "body", Name: "Bar"})
	require.NoError(t, err)
	require.Contains(t, out, `"id": "n-bar"`)
}

func TestTokensaveFilesEntitiesStatus(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)
	ctx := context.Background()

	out, err := runTokensave(ctx, root, &TokensaveParams{Op: "files"})
	require.NoError(t, err)
	require.Contains(t, out, `"path": "a.go"`)
	require.Contains(t, out, `"path": "b.go"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "files", File: "b"})
	require.NoError(t, err)
	require.Contains(t, out, `"path": "b.go"`)
	require.NotContains(t, out, `"path": "a.go"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "entities", File: "a.go"})
	require.NoError(t, err)
	require.Contains(t, out, `"name": "Foo"`)
	require.Contains(t, out, `"name": "Bar"`)
	require.NotContains(t, out, `"name": "Baz"`)

	out, err = runTokensave(ctx, root, &TokensaveParams{Op: "status"})
	require.NoError(t, err)
	require.Contains(t, out, `"nodes": 3`)
	require.Contains(t, out, `"edges": 1`)
	require.Contains(t, out, `"files": 2`)
}

func TestTokensaveBodyWithStaleGraph(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)

	db, err := sql.Open("sqlite", "file:"+filepath.Join(root, ".tokensave", "tokensave.db"))
	require.NoError(t, err)
	_, err = db.Exec(`INSERT INTO nodes (id, kind, name, qualified_name, file_path, start_line, end_line)
		VALUES ('n-stale', 'function', 'Stale', 'pkg.Stale', 'a.go', 99, 200)`)
	require.NoError(t, err)
	require.NoError(t, db.Close())

	// Source file only has 8 lines; the stale range must not panic.
	out, err := runTokensave(context.Background(), root, &TokensaveParams{Op: "body", ID: "n-stale"})
	require.NoError(t, err)
	require.Contains(t, out, `"code": ""`)
}

func TestTokensaveErrors(t *testing.T) {
	t.Parallel()
	root := tokensaveFixture(t)
	ctx := context.Background()

	_, err := runTokensave(ctx, t.TempDir(), &TokensaveParams{Op: "status"})
	require.ErrorContains(t, err, "tokensave graph not found")

	_, err = runTokensave(ctx, root, &TokensaveParams{Op: "bogus"})
	require.ErrorContains(t, err, "unknown op")

	_, err = runTokensave(ctx, root, &TokensaveParams{Op: "search"})
	require.ErrorContains(t, err, "query is required")

	_, err = runTokensave(ctx, root, &TokensaveParams{Op: "node", ID: "n-nope"})
	require.ErrorContains(t, err, "not found")
}
