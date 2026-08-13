package tools

import (
	"path/filepath"

	mem "github.com/Gentleman-Programming/engram/mem"
)

// openMemStore opens (creating if needed) the engram store under the Crush
// data directory. Each call opens a fresh handle; use-and-close keeps the
// handle's SQLite connection from lingering across tool calls. Store limits
// fall back to engram's built-in defaults via mem.Open.
func openMemStore(dataDir string) (*mem.Handle, error) {
	return mem.Open(mem.Config{
		DataDir: filepath.Join(dataDir, "engram"),
	})
}
