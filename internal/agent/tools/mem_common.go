package tools

import (
	"path/filepath"
	"sync"

	mem "github.com/Gentleman-Programming/engram/mem"
)

// memStores caches one open engram handle per data directory so repeated
// tool calls do not re-run the (migration-heavy) store.New on every use.
// Keys are the absolute data-dir paths, values are *mem.Handle. Crush runs a
// single workspace per process, so the map stays tiny.
var (
	memStoresMu sync.Mutex
	memStores   = map[string]*mem.Handle{}
)

// openMemStore returns the engram store for the Crush data directory,
// opening (and caching) it on first use. The handle is never closed on
// purpose: it lives for the process, matching engram's daemon-style reuse.
func openMemStore(dataDir string) (*mem.Handle, error) {
	dir := filepath.Join(dataDir, "engram")
	memStoresMu.Lock()
	defer memStoresMu.Unlock()
	if h, ok := memStores[dir]; ok {
		return h, nil
	}
	h, err := mem.Open(mem.Config{DataDir: dir})
	if err != nil {
		return nil, err
	}
	memStores[dir] = h
	return h, nil
}
