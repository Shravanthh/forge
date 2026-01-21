package ctx

import "sync"

// EventHandler handles UI events.
type EventHandler func(*Context)

// Context holds state and event handlers for a session.
type Context struct {
	mu         sync.RWMutex
	state      map[string]any
	persistent map[string]bool
	events     map[string]EventHandler
	Params     map[string]string
}

// New creates a new Context.
func New() *Context {
	return &Context{
		state:      make(map[string]any),
		persistent: make(map[string]bool),
		events:     make(map[string]EventHandler),
		Params:     make(map[string]string),
	}
}

// Set stores a value.
func (c *Context) Set(key string, val any) {
	c.mu.Lock()
	c.state[key] = val
	c.mu.Unlock()
}

// Get retrieves a value.
func (c *Context) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state[key]
}

// Int returns an int value or 0.
func (c *Context) Int(key string) int {
	v, _ := c.Get(key).(int)
	return v
}

// String returns a string value or "".
func (c *Context) String(key string) string {
	v, _ := c.Get(key).(string)
	return v
}

// Bool returns a bool value or false.
func (c *Context) Bool(key string) bool {
	v, _ := c.Get(key).(bool)
	return v
}

// Persist marks a key for session persistence.
func (c *Context) Persist(key string) {
	c.mu.Lock()
	c.persistent[key] = true
	c.mu.Unlock()
}

// PersistentState returns all persistent state.
func (c *Context) PersistentState() map[string]any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result := make(map[string]any)
	for k := range c.persistent {
		if v, ok := c.state[k]; ok {
			result[k] = v
		}
	}
	return result
}

// RestoreState restores persistent state.
func (c *Context) RestoreState(state map[string]any) {
	c.mu.Lock()
	for k, v := range state {
		c.state[k] = v
		c.persistent[k] = true
	}
	c.mu.Unlock()
}

// On registers an event handler.
func (c *Context) On(id string, handler EventHandler) {
	c.mu.Lock()
	c.events[id] = handler
	c.mu.Unlock()
}

// Handle executes an event handler.
func (c *Context) Handle(id string) bool {
	c.mu.RLock()
	handler, ok := c.events[id]
	c.mu.RUnlock()
	if ok {
		handler(c)
	}
	return ok
}

// HandleWithValue sets input value then executes handler.
func (c *Context) HandleWithValue(id, value string) bool {
	c.Set("_input", value)
	return c.Handle(id)
}

// InputValue returns the current input value from an event.
func (c *Context) InputValue() string { return c.String("_input") }
