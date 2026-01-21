# State Management

Forge uses server-side state management through the Context object.

## Reading State

```go
func Page(c *forge.Context) ui.UI {
    // Type-safe getters
    name := c.String("name")     // returns "" if not found
    count := c.Int("count")      // returns 0 if not found
    active := c.Bool("active")   // returns false if not found
    
    // Generic getter
    value := c.Get("key")        // returns any (nil if not found)
}
```

## Writing State

```go
func handleClick(c *forge.Context) {
    c.Set("name", "John")
    c.Set("count", 42)
    c.Set("active", true)
    c.Set("data", map[string]any{"key": "value"})
}
```

## State in Event Handlers

```go
func Page(c *forge.Context) ui.UI {
    count := c.Int("count")
    
    return ui.Div(
        ui.P(ui.T(fmt.Sprintf("Count: %d", count))),
        ui.Button(ui.T("Increment")).
            WithID("inc").
            OnClick(c, func(c *forge.Context) {
                c.Set("count", c.Int("count") + 1)
            }),
    )
}
```

## Persistent State

By default, state is lost when the WebSocket disconnects. Mark keys as persistent to survive reconnection:

```go
func Page(c *forge.Context) ui.UI {
    // Initialize and persist
    if c.Get("user_id") == nil {
        c.Set("user_id", generateID())
        c.Persist("user_id")  // survives reconnection
    }
}
```

## Route Parameters

Access URL parameters via `c.Params`:

```go
// Route: /user/:id
app.Route("/user/:id", UserPage)

func UserPage(c *forge.Context) ui.UI {
    userID := c.Params["id"]
    return ui.H1(ui.T("User: " + userID))
}
```

## Input Values

Get input values in event handlers:

```go
ui.Input().
    WithID("name-input").
    OnInput(c, func(c *forge.Context) {
        value := c.InputValue()
        c.Set("name", value)
    })
```

## State Patterns

### Toggle

```go
ui.Button(ui.T("Toggle")).OnClick(c, func(c *forge.Context) {
    c.Set("open", !c.Bool("open"))
})
```

### Counter

```go
ui.Button(ui.T("-")).OnClick(c, func(c *forge.Context) {
    c.Set("count", c.Int("count") - 1)
})
ui.Button(ui.T("+")).OnClick(c, func(c *forge.Context) {
    c.Set("count", c.Int("count") + 1)
})
```

### Form Data

```go
func handleSubmit(c *forge.Context) {
    name := c.String("form_name")
    email := c.String("form_email")
    // Process form...
    
    // Clear form
    c.Set("form_name", "")
    c.Set("form_email", "")
}
```

### List Management

```go
// Add item
func addTodo(c *forge.Context) {
    todos := c.Get("todos").([]string)
    todos = append(todos, c.String("new_todo"))
    c.Set("todos", todos)
    c.Set("new_todo", "")
}

// Remove item
func removeTodo(c *forge.Context, index int) {
    todos := c.Get("todos").([]string)
    todos = append(todos[:index], todos[index+1:]...)
    c.Set("todos", todos)
}
```

## Thread Safety

Context is thread-safe. All reads and writes are protected by a mutex.
