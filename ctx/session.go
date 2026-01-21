package ctx

import "sync"

// SessionStore is the interface for persisting session state.
// Implement this interface to use custom storage (Redis, database, etc.).
type SessionStore interface {
	// Save persists the state for a session ID.
	Save(id string, state map[string]any) error
	// Load retrieves the state for a session ID.
	Load(id string) (map[string]any, error)
}

// MemoryStore is an in-memory implementation of SessionStore.
// Suitable for development and single-instance deployments.
// For production with multiple instances, use Redis or database storage.
type MemoryStore struct {
	mu    sync.RWMutex
	store map[string]map[string]any
}

// NewMemoryStore creates a new in-memory session store.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{store: make(map[string]map[string]any)}
}

// Save stores the session state in memory.
func (m *MemoryStore) Save(id string, state map[string]any) error {
	m.mu.Lock()
	m.store[id] = state
	m.mu.Unlock()
	return nil
}

// Load retrieves the session state from memory.
func (m *MemoryStore) Load(id string) (map[string]any, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.store[id], nil
}
