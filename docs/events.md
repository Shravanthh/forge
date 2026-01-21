# Events

Forge handles events on the server via WebSocket.

## Available Events

| Event | Method | Use Case |
|-------|--------|----------|
| Click | `OnClick` | Buttons, links, any clickable |
| Input | `OnInput` | Text input (fires on every keystroke) |
| Change | `OnChange` | Select, checkbox, radio |
| Submit | `OnSubmit` | Form submission |
| Keydown | `OnKeydown` | Keyboard events (Enter key) |
| Scroll | `OnScroll` | Scroll position tracking |

## Click Events

```go
ui.Button(ui.T("Click me")).
    WithID("my-btn").
    OnClick(c, func(c *forge.Context) {
        c.Set("clicked", true)
    })
```

## Input Events

Fires on every keystroke:

```go
ui.Input().
    WithID("search").
    WithAttr("type", "text").
    WithAttr("placeholder", "Search...").
    OnInput(c, func(c *forge.Context) {
        query := c.InputValue()
        c.Set("search_query", query)
    })
```

## Change Events

For select, checkbox, and radio:

```go
// Checkbox
ui.Input().
    WithID("agree").
    WithAttr("type", "checkbox").
    OnChange(c, func(c *forge.Context) {
        checked := c.InputValue() == "true"
        c.Set("agreed", checked)
    })

// Select (use El for select element)
ui.El("select",
    ui.El("option", ui.T("Option 1")).WithAttr("value", "1"),
    ui.El("option", ui.T("Option 2")).WithAttr("value", "2"),
).WithID("choice").OnChange(c, func(c *forge.Context) {
    selected := c.InputValue()
    c.Set("choice", selected)
})
```

## Submit Events

```go
ui.Form(
    ui.Input().WithID("email").WithAttr("type", "email").
        OnInput(c, func(c *forge.Context) {
            c.Set("email", c.InputValue())
        }),
    ui.Button(ui.T("Submit")).WithAttr("type", "submit"),
).WithID("my-form").OnSubmit(c, func(c *forge.Context) {
    email := c.String("email")
    // Process form...
})
```

## Keydown Events

Triggers on Enter key:

```go
ui.Input().
    WithID("todo-input").
    WithAttr("placeholder", "Press Enter to add").
    OnKeydown(c, func(c *forge.Context) {
        text := c.InputValue()
        if text != "" {
            // Add todo
            c.Set("input", "")
        }
    })
```

## Event Handler Patterns

### Toggle

```go
OnClick(c, func(c *forge.Context) {
    c.Set("open", !c.Bool("open"))
})
```

### Increment/Decrement

```go
// Decrement
OnClick(c, func(c *forge.Context) {
    c.Set("count", c.Int("count") - 1)
})

// Increment
OnClick(c, func(c *forge.Context) {
    c.Set("count", c.Int("count") + 1)
})
```

### Add to List

```go
OnClick(c, func(c *forge.Context) {
    items := c.Get("items").([]string)
    newItem := c.String("new_item")
    items = append(items, newItem)
    c.Set("items", items)
    c.Set("new_item", "")
})
```

### Remove from List

```go
// Capture index in closure
idx := i
OnClick(c, func(c *forge.Context) {
    items := c.Get("items").([]string)
    items = append(items[:idx], items[idx+1:]...)
    c.Set("items", items)
})
```

## Element IDs

Every element with an event handler needs a unique ID:

```go
// Good - unique IDs
ui.Button(ui.T("Save")).WithID("save-btn").OnClick(c, handleSave)
ui.Button(ui.T("Cancel")).WithID("cancel-btn").OnClick(c, handleCancel)

// For lists, include index
for i, item := range items {
    ui.Button(ui.T("Delete")).
        WithID(fmt.Sprintf("delete-%d", i)).
        OnClick(c, deleteHandler(i))
}
```

## Event Debouncing

For scroll events, debouncing is built-in (100ms). For other events, implement in handler:

```go
var lastCall time.Time

OnInput(c, func(c *forge.Context) {
    if time.Since(lastCall) < 300*time.Millisecond {
        return
    }
    lastCall = time.Now()
    // Handle input...
})
```
