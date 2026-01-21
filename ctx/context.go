package ctx

import "sync"

// EventHandler is a function that handles UI events.
// It receives the Context to read/write state.
//
//	func handleClick(c *ctx.Context) {
//	    c.Set("count", c.Int("count") + 1)
//	}
type EventHandler func(*Context)

// Context holds the state, event handlers, and route parameters for a session.
// It is the central state container passed to all page functions and event handlers.
//
// # Reading State
//
//	name := c.String("name")
//	count := c.Int("count")
//	active := c.Bool("active")
//	value := c.Get("key")  // returns any
//
// # Writing State
//
//	c.Set("name", "John")
//	c.Set("count", 42)
//
// # Persistence
//
//	c.Persist("user_id")  // survives reconnection
//
// # Route Parameters
//
//	// For route "/user/:id"
//	userID := c.Params["id"]
type Context struct {
	mu         sync.RWMutex
	state      map[string]any
	persistent map[string]bool
	events     map[string]EventHandler
	Params     map[string]string // Route parameters (e.g., :id)
}

// New creates a new empty Context.
func New() *Context {
	return &Context{
		state:      make(map[string]any),
		persistent: make(map[string]bool),
		events:     make(map[string]EventHandler),
		Params:     make(map[string]string),
	}
}

// Set stores a value in the context state.
func (c *Context) Set(key string, val any) {
	c.mu.Lock()
	c.state[key] = val
	c.mu.Unlock()
}

// Get retrieves a value from the context state.
// Returns nil if the key doesn't exist.
func (c *Context) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.state[key]
}

// Int returns an int value or 0 if not found or wrong type.
func (c *Context) Int(key string) int {
	v, _ := c.Get(key).(int)
	return v
}

// String returns a string value or "" if not found or wrong type.
func (c *Context) String(key string) string {
	v, _ := c.Get(key).(string)
	return v
}

// Bool returns a bool value or false if not found or wrong type.
func (c *Context) Bool(key string) bool {
	v, _ := c.Get(key).(bool)
	return v
}

// Persist marks a key for session persistence.
// Persistent values survive WebSocket reconnection.
func (c *Context) Persist(key string) {
	c.mu.Lock()
	c.persistent[key] = true
	c.mu.Unlock()
}

// PersistentState returns all persistent state as a map.
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

// RestoreState restores previously persisted state.
func (c *Context) RestoreState(state map[string]any) {
	c.mu.Lock()
	for k, v := range state {
		c.state[k] = v
		c.persistent[k] = true
	}
	c.mu.Unlock()
}

// On registers an event handler with the given ID.
func (c *Context) On(id string, handler EventHandler) {
	c.mu.Lock()
	c.events[id] = handler
	c.mu.Unlock()
}

// Handle executes the event handler with the given ID.
// Returns true if the handler was found and executed.
func (c *Context) Handle(id string) bool {
	c.mu.RLock()
	handler, ok := c.events[id]
	c.mu.RUnlock()
	if ok {
		handler(c)
	}
	return ok
}

// HandleWithValue sets the input value then executes the handler.
// Used for input events where the value needs to be passed.
func (c *Context) HandleWithValue(id, value string) bool {
	c.Set("_input", value)
	return c.Handle(id)
}

// InputValue returns the current input value from an event.
// Call this inside an OnInput or OnKeydown handler.
func (c *Context) InputValue() string { return c.String("_input") }
