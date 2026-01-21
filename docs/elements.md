# UI Elements

Forge provides a declarative DSL for building UI in Go.

## Basic Elements

### Layout Elements

```go
ui.Div(children...)      // <div>
ui.Span(children...)     // <span>
ui.Main(children...)     // <main>
ui.Section(children...)  // <section>
ui.Article(children...)  // <article>
ui.Header(children...)   // <header>
ui.Footer(children...)   // <footer>
ui.Nav(children...)      // <nav>
```

### Text Elements

```go
ui.P(children...)   // <p>
ui.H1(children...)  // <h1>
ui.H2(children...)  // <h2>
ui.H3(children...)  // <h3>
ui.H4(children...)  // <h4>
ui.T("text")        // text node
```

### Interactive Elements

```go
ui.Button(children...)  // <button>
ui.A(children...)       // <a>
ui.Form(children...)    // <form>
ui.Label(children...)   // <label>
ui.Input()              // <input>
```

### List Elements

```go
ui.Ul(children...)  // <ul>
ui.Li(children...)  // <li>
```

### Table Elements

```go
ui.Tr(children...)  // <tr>
ui.Th(children...)  // <th>
ui.Td(children...)  // <td>
```

### Self-Closing Elements

```go
ui.Input()  // <input />
ui.Img()    // <img />
ui.Br()     // <br />
ui.Hr()     // <hr />
```

### Custom Elements

```go
ui.El("custom-tag", children...)
```

## Attributes

Chain methods to add attributes:

```go
ui.Div(
    ui.T("Content"),
).WithID("main").
  WithClass("container").
  WithStyle("color: red").
  WithAttr("data-value", "123")
```

### Available Methods

| Method | Description |
|--------|-------------|
| `WithID(id)` | Set element ID |
| `WithClass(class)` | Set CSS class |
| `WithStyle(style)` | Set inline style |
| `WithS(style)` | Set style using Style builder |
| `WithAttr(key, value)` | Set custom attribute |
| `WithChildren(children...)` | Append children |

## Composition

Create reusable components as functions:

```go
func Card(title string, content ui.UI) ui.UI {
    return ui.Div(
        ui.H3(ui.T(title)).WithClass("card-title"),
        ui.Div(content).WithClass("card-body"),
    ).WithClass("card")
}

// Usage
Card("Welcome", ui.P(ui.T("Hello!")))
```

## Conditional Rendering

```go
func Page(c *forge.Context) ui.UI {
    isLoggedIn := c.Bool("logged_in")
    
    var content ui.UI
    if isLoggedIn {
        content = ui.T("Welcome back!")
    } else {
        content = ui.T("Please log in")
    }
    
    return ui.Div(content)
}
```

## List Rendering

```go
func TodoList(c *forge.Context) ui.UI {
    todos := []string{"Buy milk", "Walk dog", "Code"}
    
    var items []ui.UI
    for i, todo := range todos {
        items = append(items, 
            ui.Li(ui.T(todo)).WithID(fmt.Sprintf("todo-%d", i)),
        )
    }
    
    return ui.Ul(items...)
}
```

## Raw HTML

Use with caution (XSS risk):

```go
ui.Raw{HTML: "<strong>Bold text</strong>"}
```
