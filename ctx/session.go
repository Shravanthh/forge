package ctx

import "sync"

// SessionStore persists state across connections.
type SessionStore interface {
	Save(id string, state map[string]any) error
	Load(id string) (map[string]any, error)
}

// MemoryStore is an in-memory session store.
type MemoryStore struct {
	mu    sync.RWMutex
	store map[string]map[string]any
}

// NewMemoryStore creates a new in-memory store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{store: make(map[string]map[string]any)}
}

func (m *MemoryStore) Save(id string, state map[string]any) error {
	m.mu.Lock()
	m.store[id] = state
	m.mu.Unlock()
	return nil
}

func (m *MemoryStore) Load(id string) (map[string]any, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.store[id], nil
}
