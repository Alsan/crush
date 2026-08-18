package tools

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"charm.land/fantasy"
	_ "modernc.org/sqlite"
)

// TokensaveToolName queries the tokensave code-graph database directly.
const TokensaveToolName = "tokensave"

//go:embed tokensave.md
var tokensaveDescription string

const (
	tokensaveDBRelPath        = ".tokensave/tokensave.db"
	tokensaveQueryTimeout     = 10 * time.Second
	tokensaveDefaultLimit     = 20
	tokensaveMaxLimit         = 200
	tokensaveDefaultImpactDep = 3
	tokensaveMaxDepth         = 10
	tokensaveMaxBodyLines     = 400
	tokensaveMaxBodyBytes     = 2 << 20
	tokensaveMaxVisited       = 500
)

// tokensaveOps are the operations the tool accepts.
var tokensaveOps = []string{
	"status", "files", "entities", "search", "find",
	"node", "body", "callers", "callees", "impact",
	"complexity", "test_map", "deps",
}

// TokensaveParams are the inputs for the tokensave tool.
type TokensaveParams struct {
	Op        string `json:"op" description:"Operation: status|files|entities|search|find|node|body|callers|callees|impact|complexity|test_map|deps"`
	Root      string `json:"root,omitempty" description:"Project root containing .tokensave/tokensave.db (defaults to working dir)"`
	ID        string `json:"id,omitempty" description:"Node ID from a previous search/find result"`
	Name      string `json:"name,omitempty" description:"Exact symbol name (for find/body)"`
	Query     string `json:"query,omitempty" description:"Search terms (for search)"`
	File      string `json:"file,omitempty" description:"Project-relative file path (for entities/files)"`
	Kinds     string `json:"kinds,omitempty" description:"Comma-separated node kinds filter"`
	EdgeKinds string `json:"edge_kinds,omitempty" description:"Comma-separated edge kinds for graph ops (default calls)"`
	MaxDepth  int    `json:"max_depth,omitempty" description:"Traversal depth for callers/callees/impact"`
	Limit     int    `json:"limit,omitempty" description:"Max results (default 20)"`
}

type tokensaveSymbol struct {
	ID        string `json:"id"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Signature string `json:"signature,omitempty"`
}

type tokensaveGraphSymbol struct {
	tokensaveSymbol
	Depth int `json:"depth"`
	Edge  int `json:"edge_line,omitempty"`
}

type tokensaveNodeDetail struct {
	ID                  string `json:"id"`
	Kind                string `json:"kind"`
	Name                string `json:"name"`
	QualifiedName       string `json:"qualified_name"`
	File                string `json:"file"`
	StartLine           int    `json:"start_line"`
	EndLine             int    `json:"end_line"`
	Signature           string `json:"signature,omitempty"`
	Docstring           string `json:"docstring,omitempty"`
	Visibility          string `json:"visibility"`
	Async               bool   `json:"async"`
	CognitiveComplexity int    `json:"cognitive_complexity"`
	IncomingEdges       int    `json:"incoming_edges"`
	OutgoingEdges       int    `json:"outgoing_edges"`
}

type tokensaveFileRow struct {
	Path      string `json:"path"`
	NodeCount int    `json:"node_count"`
	Size      int    `json:"size"`
	Kind      string `json:"kind"`
}

type tokensaveStatus struct {
	Root  string `json:"root"`
	Nodes int    `json:"nodes"`
	Edges int    `json:"edges"`
	Files int    `json:"files"`
	DBMB  int    `json:"db_mb"`
}

func NewTokensaveTool(workingDir string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		TokensaveToolName,
		tokensaveDescription,
		func(ctx context.Context, params TokensaveParams, _ fantasy.ToolCall) (fantasy.ToolResponse, error) {
			root := params.Root
			if root == "" {
				root = workingDir
			}
			absRoot, err := filepath.Abs(root)
			if err != nil {
				return fantasy.NewTextErrorResponse(err.Error()), nil
			}
			out, err := runTokensave(ctx, absRoot, &params)
			if err != nil {
				return fantasy.NewTextErrorResponse(err.Error()), nil
			}
			return fantasy.NewTextResponse(out), nil
		},
	)
}

// runTokensave executes one tokensave operation against the graph DB in
// root and returns JSON output.
func runTokensave(ctx context.Context, root string, p *TokensaveParams) (string, error) {
	if p.Op == "" {
		return "", fmt.Errorf("op is required (one of: %s)", strings.Join(tokensaveOps, ", "))
	}
	if !tokensaveValidOp(p.Op) {
		return "", fmt.Errorf("unknown op %q (allowed: %s)", p.Op, strings.Join(tokensaveOps, ", "))
	}
	dbPath := filepath.Join(root, tokensaveDBRelPath)
	if _, err := os.Stat(dbPath); err != nil {
		return "", fmt.Errorf("tokensave graph not found at %s; initialize the tokensave MCP server for this repository first", dbPath)
	}
	dsnParams := url.Values{}
	dsnParams.Set("mode", "ro")
	dsnParams.Set("_txlock", "immediate")
	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s?%s", dbPath, dsnParams.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to open graph database: %w", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	qctx, cancel := context.WithTimeout(ctx, tokensaveQueryTimeout)
	defer cancel()

	g := &tokensaveGraph{db: db, root: root, limit: tokensaveClampLimit(p.Limit)}
	switch p.Op {
	case "status":
		return g.status(qctx)
	case "files":
		return g.files(qctx, p)
	case "entities":
		if p.File == "" {
			return "", fmt.Errorf("file is required for entities")
		}
		return g.entities(qctx, p)
	case "search":
		if p.Query == "" {
			return "", fmt.Errorf("query is required for search")
		}
		return g.search(qctx, p)
	case "find":
		if p.Name == "" {
			return "", fmt.Errorf("name is required for find")
		}
		return g.find(qctx, p)
	case "node":
		if p.ID == "" {
			return "", fmt.Errorf("id is required for node")
		}
		return g.node(qctx, p.ID)
	case "body":
		if p.ID == "" && p.Name == "" {
			return "", fmt.Errorf("id or name is required for body")
		}
		return g.body(qctx, p)
	case "callers", "callees":
		if p.ID == "" {
			return "", fmt.Errorf("id is required for %s", p.Op)
		}
		return g.traverseNeighbors(qctx, p, p.Op == "callers", 1)
	case "impact":
		if p.ID == "" {
			return "", fmt.Errorf("id is required for impact")
		}
		return g.traverseNeighbors(qctx, p, true, tokensaveDefaultImpactDep)
	case "complexity":
		return g.complexity(qctx, p)
	case "test_map":
		return g.testMap(qctx, p)
	case "deps":
		return g.deps(qctx, p)
	}
	return "", fmt.Errorf("unhandled op %q", p.Op)
}

type tokensaveGraph struct {
	db    *sql.DB
	root  string
	limit int
}

func (g *tokensaveGraph) status(ctx context.Context) (string, error) {
	var s tokensaveStatus
	s.Root = g.root
	for q, dst := range map[string]*int{
		"SELECT count(*) FROM nodes": &s.Nodes,
		"SELECT count(*) FROM edges": &s.Edges,
		"SELECT count(*) FROM files": &s.Files,
	} {
		if err := g.db.QueryRowContext(ctx, q).Scan(dst); err != nil {
			return "", fmt.Errorf("failed to read graph stats: %w", err)
		}
	}
	if fi, err := os.Stat(filepath.Join(g.root, tokensaveDBRelPath)); err == nil {
		s.DBMB = int(fi.Size() >> 20)
	}
	return tokensaveJSON(s)
}

func (g *tokensaveGraph) files(ctx context.Context, p *TokensaveParams) (string, error) {
	q := "SELECT path, node_count, size, kind FROM files"
	args := []any{}
	if p.File != "" {
		q += " WHERE path LIKE ?"
		args = append(args, p.File+"%")
	}
	q += " ORDER BY path LIMIT ?"
	args = append(args, g.limit)
	rows, err := g.db.QueryContext(ctx, q, args...)
	if err != nil {
		return "", fmt.Errorf("failed to list files: %w", err)
	}
	defer rows.Close()
	var out []tokensaveFileRow
	for rows.Next() {
		var r tokensaveFileRow
		if err := rows.Scan(&r.Path, &r.NodeCount, &r.Size, &r.Kind); err != nil {
			return "", fmt.Errorf("failed to scan file row: %w", err)
		}
		out = append(out, r)
	}
	return tokensaveJSON(out)
}

func (g *tokensaveGraph) entities(ctx context.Context, p *TokensaveParams) (string, error) {
	q := "SELECT id, kind, name, file_path, start_line, COALESCE(signature,'') FROM nodes WHERE file_path = ?"
	args := []any{p.File}
	if kinds := tokensaveKinds(p.Kinds); len(kinds) > 0 {
		q += " AND kind IN (" + strings.Repeat("?,", len(kinds)-1) + "?)"
		for _, k := range kinds {
			args = append(args, k)
		}
	}
	q += " ORDER BY start_line LIMIT ?"
	args = append(args, g.limit)
	rows, err := g.db.QueryContext(ctx, q, args...)
	if err != nil {
		return "", fmt.Errorf("failed to list entities: %w", err)
	}
	defer rows.Close()
	var out []tokensaveSymbol
	for rows.Next() {
		var s tokensaveSymbol
		if err := rows.Scan(&s.ID, &s.Kind, &s.Name, &s.File, &s.Line, &s.Signature); err != nil {
			return "", fmt.Errorf("failed to scan entity row: %w", err)
		}
		out = append(out, s)
	}
	return tokensaveJSON(out)
}

func (g *tokensaveGraph) search(ctx context.Context, p *TokensaveParams) (string, error) {
	base := `SELECT n.id, n.kind, n.name, n.file_path, n.start_line, COALESCE(n.signature,'')
		FROM nodes_fts f JOIN nodes n ON n.rowid = f.rowid
		WHERE nodes_fts MATCH ?`
	kinds := tokensaveKinds(p.Kinds)
	if len(kinds) > 0 {
		base += " AND n.kind IN (" + strings.Repeat("?,", len(kinds)-1) + "?)"
	}
	base += " ORDER BY bm25(nodes_fts) LIMIT ?"
	args := []any{p.Query}
	for _, k := range kinds {
		args = append(args, k)
	}
	args = append(args, g.limit)
	rows, err := g.db.QueryContext(ctx, base, args...)
	if err != nil {
		// Raw FTS5 syntax can be rejected; retry with each term quoted.
		quoted := make([]string, 0, 4)
		for _, term := range strings.Fields(p.Query) {
			quoted = append(quoted, `"`+strings.ReplaceAll(term, `"`, "")+`"`)
		}
		if len(quoted) == 0 {
			return "", fmt.Errorf("failed to search graph: %w", err)
		}
		args[0] = strings.Join(quoted, " ")
		rows, err = g.db.QueryContext(ctx, base, args...)
		if err != nil {
			return "", fmt.Errorf("failed to search graph: %w", err)
		}
	}
	defer rows.Close()
	out, err := tokensaveScanSymbols(rows)
	if err != nil {
		return "", err
	}
	return tokensaveJSON(out)
}

func (g *tokensaveGraph) find(ctx context.Context, p *TokensaveParams) (string, error) {
	q := `SELECT id, kind, name, file_path, start_line, COALESCE(signature,'')
		FROM nodes WHERE name = ?`
	args := []any{p.Name}
	if kinds := tokensaveKinds(p.Kinds); len(kinds) > 0 {
		q += " AND kind IN (" + strings.Repeat("?,", len(kinds)-1) + "?)"
		for _, k := range kinds {
			args = append(args, k)
		}
	}
	q += " ORDER BY length(qualified_name), qualified_name LIMIT ?"
	args = append(args, g.limit)
	rows, err := g.db.QueryContext(ctx, q, args...)
	if err != nil {
		return "", fmt.Errorf("failed to find symbol: %w", err)
	}
	defer rows.Close()
	out, err := tokensaveScanSymbols(rows)
	if err != nil {
		return "", err
	}
	return tokensaveJSON(out)
}

func (g *tokensaveGraph) node(ctx context.Context, id string) (string, error) {
	var n tokensaveNodeDetail
	var sig, doc sql.NullString
	var isAsync int
	err := g.db.QueryRowContext(ctx, `SELECT id, kind, name, qualified_name, file_path,
			start_line, end_line, signature, docstring, visibility, is_async, cognitive_complexity
		FROM nodes WHERE id = ?`, id).
		Scan(&n.ID, &n.Kind, &n.Name, &n.QualifiedName, &n.File,
			&n.StartLine, &n.EndLine, &sig, &doc, &n.Visibility, &isAsync, &n.CognitiveComplexity)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("node %q not found; use search or find to get a valid id", id)
	}
	if err != nil {
		return "", fmt.Errorf("failed to read node: %w", err)
	}
	n.Signature = sig.String
	n.Docstring = doc.String
	n.Async = isAsync != 0
	_ = g.db.QueryRowContext(ctx,
		"SELECT count(*) FROM edges WHERE target = ?", id).Scan(&n.IncomingEdges)
	_ = g.db.QueryRowContext(ctx,
		"SELECT count(*) FROM edges WHERE source = ?", id).Scan(&n.OutgoingEdges)
	return tokensaveJSON(n)
}

func (g *tokensaveGraph) body(ctx context.Context, p *TokensaveParams) (string, error) {
	id := p.ID
	if id == "" {
		rows, err := g.db.QueryContext(ctx,
			`SELECT id FROM nodes WHERE name = ?
			ORDER BY length(qualified_name), qualified_name LIMIT 1`, p.Name)
		if err != nil {
			return "", fmt.Errorf("failed to find symbol: %w", err)
		}
		defer rows.Close()
		hasRow := rows.Next()
		if hasRow {
			if err := rows.Scan(&id); err != nil {
				rows.Close()
				return "", fmt.Errorf("failed to scan symbol id: %w", err)
			}
		}
		// Close before issuing the next query: the pool has one connection.
		if err := rows.Close(); err != nil {
			return "", fmt.Errorf("failed to find symbol: %w", err)
		}
		if !hasRow {
			return "", fmt.Errorf("symbol %q not found", p.Name)
		}
	}
	var filePath string
	var start, end int
	err := g.db.QueryRowContext(ctx,
		"SELECT file_path, start_line, end_line FROM nodes WHERE id = ?", id).
		Scan(&filePath, &start, &end)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("node %q not found", id)
	}
	if err != nil {
		return "", fmt.Errorf("failed to read node location: %w", err)
	}
	src, err := tokensaveReadSource(g.root, filePath)
	if err != nil {
		return "", err
	}
	lines := strings.Split(src, "\n")
	if start < 1 {
		start = 1
	}
	if end > len(lines) {
		end = len(lines)
	}
	// The graph can be stale relative to the working tree; clamp so the
	// slice expression below can never go out of bounds.
	if start > len(lines) {
		start = len(lines) + 1
	}
	if end < start-1 {
		end = start - 1
	}
	if end-start+1 > tokensaveMaxBodyLines {
		end = start + tokensaveMaxBodyLines - 1
	}
	body := struct {
		ID   string `json:"id"`
		File string `json:"file"`
		From int    `json:"from_line"`
		To   int    `json:"to_line"`
		Code string `json:"code"`
	}{ID: id, File: filePath, From: start, To: end, Code: strings.Join(lines[start-1:end], "\n")}
	return tokensaveJSON(body)
}

// traverseNeighbors walks edges from the node either backwards (incoming,
// callers/impact) or forwards (outgoing, callees), collecting symbols with
// their BFS depth, bounded by maxDepth and a visited cap.
func (g *tokensaveGraph) traverseNeighbors(ctx context.Context, p *TokensaveParams, incoming bool, defaultDepth int) (string, error) {
	depth := p.MaxDepth
	if depth <= 0 {
		depth = defaultDepth
	}
	if depth > tokensaveMaxDepth {
		depth = tokensaveMaxDepth
	}
	edgeKinds := tokensaveKinds(p.EdgeKinds)
	if len(edgeKinds) == 0 {
		edgeKinds = []string{"calls"}
	}
	edgeFilter := " AND e.kind IN (" + strings.Repeat("?,", len(edgeKinds)-1) + "?)"
	kindArgs := make([]any, len(edgeKinds))
	for i, k := range edgeKinds {
		kindArgs[i] = k
	}

	type frontier struct {
		id    string
		depth int
	}
	queue := []frontier{{id: p.ID, depth: 0}}
	visited := map[string]bool{p.ID: true}
	var out []tokensaveGraphSymbol
	for len(queue) > 0 && len(out) < g.limit && len(visited) < tokensaveMaxVisited {
		cur := queue[0]
		queue = queue[1:]
		if cur.depth >= depth {
			continue
		}
		q := `SELECT src.id, src.kind, src.name, src.file_path, src.start_line,
				COALESCE(src.signature,''), e.line
			FROM edges e
			JOIN nodes src ON src.id = e.` + tokensaveEdgeCol(incoming) + `
			WHERE e.` + tokensaveEdgeCol(!incoming) + ` = ?` + edgeFilter + `
			ORDER BY src.name LIMIT ?`
		args := append([]any{cur.id}, kindArgs...)
		args = append(args, tokensaveMaxVisited)
		rows, err := g.db.QueryContext(ctx, q, args...)
		if err != nil {
			return "", fmt.Errorf("failed to traverse graph: %w", err)
		}
		for rows.Next() {
			var s tokensaveGraphSymbol
			var edgeLine sql.NullInt64
			if err := rows.Scan(&s.ID, &s.Kind, &s.Name, &s.File, &s.Line, &s.Signature, &edgeLine); err != nil {
				rows.Close()
				return "", fmt.Errorf("failed to scan graph row: %w", err)
			}
			if visited[s.ID] {
				continue
			}
			visited[s.ID] = true
			s.Depth = cur.depth + 1
			s.Edge = int(edgeLine.Int64)
			out = append(out, s)
			queue = append(queue, frontier{id: s.ID, depth: s.Depth})
			if len(out) >= g.limit {
				break
			}
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return "", fmt.Errorf("failed to traverse graph: %w", err)
		}
	}
	return tokensaveJSON(out)
}

func tokensaveEdgeCol(fromSource bool) string {
	if fromSource {
		return "source"
	}
	return "target"
}

func tokensaveScanSymbols(rows *sql.Rows) ([]tokensaveSymbol, error) {
	defer rows.Close()
	var out []tokensaveSymbol
	for rows.Next() {
		var s tokensaveSymbol
		if err := rows.Scan(&s.ID, &s.Kind, &s.Name, &s.File, &s.Line, &s.Signature); err != nil {
			return nil, fmt.Errorf("failed to scan symbol row: %w", err)
		}
		out = append(out, s)
	}
	return out, rows.Err()
}

func tokensaveReadSource(root, rel string) (string, error) {
	full := filepath.Join(root, filepath.FromSlash(rel))
	clean, err := filepath.Abs(full)
	if err != nil {
		return "", fmt.Errorf("failed to resolve %s: %w", rel, err)
	}
	rootAbs, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("failed to resolve root: %w", err)
	}
	if clean != rootAbs && !strings.HasPrefix(clean, rootAbs+string(filepath.Separator)) {
		return "", fmt.Errorf("refusing to read %s: outside project root", rel)
	}
	fi, err := os.Stat(clean)
	if err != nil {
		return "", fmt.Errorf("source file %s not found: %w", rel, err)
	}
	if fi.Size() > tokensaveMaxBodyBytes {
		return "", fmt.Errorf("source file %s too large (%d bytes)", rel, fi.Size())
	}
	b, err := os.ReadFile(clean)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", rel, err)
	}
	return string(b), nil
}

func (g *tokensaveGraph) complexity(ctx context.Context, p *TokensaveParams) (string, error) {
	q := `SELECT id, kind, name, file_path, start_line, COALESCE(signature,''),
	      cognitive_complexity, branches, loops, max_nesting
	      FROM nodes WHERE kind IN ('function','method')`
	args := []any{}
	if p.File != "" {
		q += " AND file_path LIKE ?"
		args = append(args, p.File+"%")
	}
	q += " ORDER BY cognitive_complexity DESC LIMIT ?"
	args = append(args, g.limit)
	rows, err := g.db.QueryContext(ctx, q, args...)
	if err != nil {
		return "", fmt.Errorf("failed to query complexity: %w", err)
	}
	defer rows.Close()
	type complexityRow struct {
		tokensaveSymbol
		CognitiveComplexity int `json:"cognitive_complexity"`
		Branches            int `json:"branches"`
		Loops               int `json:"loops"`
		MaxNesting          int `json:"max_nesting"`
	}
	var out []complexityRow
	for rows.Next() {
		var r complexityRow
		if err := rows.Scan(&r.ID, &r.Kind, &r.Name, &r.File, &r.Line, &r.Signature,
			&r.CognitiveComplexity, &r.Branches, &r.Loops, &r.MaxNesting); err != nil {
			return "", fmt.Errorf("failed to scan complexity row: %w", err)
		}
		out = append(out, r)
	}
	return tokensaveJSON(out)
}

type testMapEntry struct {
	TestFile  string   `json:"test_file"`
	TestFuncs []string `json:"test_functions"`
	Sources   []string `json:"source_files"`
}

func (g *tokensaveGraph) testMap(ctx context.Context, p *TokensaveParams) (string, error) {
	// Find test files by path pattern.
	testQ := `SELECT path FROM files WHERE (path LIKE '%_test.go' OR path LIKE '%_test.py'
	          OR path LIKE '%test_%' OR path LIKE '%/tests/%' OR path LIKE '%/test/%'
	          OR path LIKE '%spec_%' OR path LIKE '%_spec.go' OR path LIKE '%_spec.py')`
	args := []any{}
	if p.File != "" {
		testQ += " AND path LIKE ?"
		args = append(args, p.File+"%")
	}
	testQ += " ORDER BY path LIMIT ?"
	args = append(args, g.limit)
	testRows, err := g.db.QueryContext(ctx, testQ, args...)
	if err != nil {
		return "", fmt.Errorf("failed to query test files: %w", err)
	}
	defer testRows.Close()

	var testFiles []string
	for testRows.Next() {
		var path string
		if err := testRows.Scan(&path); err != nil {
			return "", fmt.Errorf("failed to scan test file: %w", err)
		}
		testFiles = append(testFiles, path)
	}

	// For each test file, find the functions it calls in other files.
	var out []testMapEntry
	for _, tf := range testFiles {
		entry := testMapEntry{TestFile: tf}
		// Get test functions in this file.
		funcQ := `SELECT name FROM nodes WHERE file_path = ? AND kind IN ('function','method') ORDER BY start_line`
		funcRows, err := g.db.QueryContext(ctx, funcQ, tf)
		if err != nil {
			continue
		}
		for funcRows.Next() {
			var name string
			if err := funcRows.Scan(&name); err != nil {
				continue
			}
			entry.TestFuncs = append(entry.TestFuncs, name)
		}
		funcRows.Close()

		// Find source files this test file calls into.
		srcQ := `SELECT DISTINCT n2.file_path FROM edges e
		         JOIN nodes n1 ON e.source = n1.id
		         JOIN nodes n2 ON e.target = n2.id
		         WHERE n1.file_path = ? AND n2.file_path != ? AND e.kind = 'calls'
		         AND n2.file_path NOT LIKE '%_test.go' AND n2.file_path NOT LIKE '%_test.py'
		         ORDER BY n2.file_path`
		srcRows, err := g.db.QueryContext(ctx, srcQ, tf, tf)
		if err != nil {
			continue
		}
		for srcRows.Next() {
			var path string
			if err := srcRows.Scan(&path); err != nil {
				continue
			}
			entry.Sources = append(entry.Sources, path)
		}
		srcRows.Close()

		if len(entry.TestFuncs) > 0 {
			out = append(out, entry)
		}
	}
	return tokensaveJSON(out)
}

type depsEntry struct {
	File   string   `json:"file"`
	Uses   []string `json:"uses"`   // files this file depends on
	UsedBy []string `json:"used_by"` // files that depend on this file
}

func (g *tokensaveGraph) deps(ctx context.Context, p *TokensaveParams) (string, error) {
	if p.File == "" {
		return "", fmt.Errorf("file is required for deps")
	}
	// Files that p.File uses (outgoing edges to different files).
	outQ := `SELECT DISTINCT n2.file_path FROM edges e
	         JOIN nodes n1 ON e.source = n1.id
	         JOIN nodes n2 ON e.target = n2.id
	         WHERE n1.file_path = ? AND n2.file_path != ? AND e.kind IN ('calls','uses')
	         ORDER BY n2.file_path`
	rows, err := g.db.QueryContext(ctx, outQ, p.File, p.File)
	if err != nil {
		return "", fmt.Errorf("failed to query deps: %w", err)
	}
	defer rows.Close()
	entry := depsEntry{File: p.File}
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			continue
		}
		entry.Uses = append(entry.Uses, path)
	}
	rows.Close()

	// Files that depend on p.File (incoming edges from different files).
	inQ := `SELECT DISTINCT n1.file_path FROM edges e
	        JOIN nodes n1 ON e.source = n1.id
	        JOIN nodes n2 ON e.target = n2.id
	        WHERE n2.file_path = ? AND n1.file_path != ? AND e.kind IN ('calls','uses')
	        ORDER BY n1.file_path`
	rows2, err := g.db.QueryContext(ctx, inQ, p.File, p.File)
	if err != nil {
		return "", fmt.Errorf("failed to query dependents: %w", err)
	}
	defer rows2.Close()
	for rows2.Next() {
		var path string
		if err := rows2.Scan(&path); err != nil {
			continue
		}
		entry.UsedBy = append(entry.UsedBy, path)
	}
	return tokensaveJSON(entry)
}

func tokensaveValidOp(op string) bool {
	for _, o := range tokensaveOps {
		if op == o {
			return true
		}
	}
	return false
}

func tokensaveClampLimit(limit int) int {
	if limit <= 0 {
		return tokensaveDefaultLimit
	}
	if limit > tokensaveMaxLimit {
		return tokensaveMaxLimit
	}
	return limit
}

func tokensaveKinds(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var kinds []string
	for _, k := range strings.Split(s, ",") {
		if k = strings.TrimSpace(k); k != "" {
			kinds = append(kinds, k)
		}
	}
	return kinds
}

func tokensaveJSON(v any) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to encode result: %w", err)
	}
	return string(b), nil
}
