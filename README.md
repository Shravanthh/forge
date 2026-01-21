# Forge

A Go-native UI framework for building modern web applications in pure Go.

## Features

- **Pure Go** - Write UI entirely in Go, no JavaScript required
- **Server-Driven** - State lives on server, UI updates via WebSocket
- **Live Updates** - Real-time DOM patching without page reloads
- **Tailwind CSS** - Built-in Tailwind support
- **Type Safe** - Full Go type safety for your UI code
- **Fast** - 424KB WASM client, sub-millisecond updates

## Installation

```bash
go get forge
```

## Quick Start

```go
package main

import (
    "forge/ctx"
    "forge/server"
    "forge/ui"
)

func init() {
    ui.EnableTailwind()
}

func main() {
    app := server.New()
    app.Route("/", HomePage)
    app.Run(":3000")
}

func HomePage(c *ctx.Context) ui.UI {
    count := c.Int("count")
    
    return ui.Div(
        ui.H1(ui.T("Counter")).WithClass("text-2xl font-bold"),
        ui.P(ui.T("Count: " + itoa(count))),
        ui.Button(ui.T("Increment")).
            WithClass("px-4 py-2 bg-blue-500 text-white rounded").
            WithID("inc").
            OnClick(c, func(c *ctx.Context) {
                c.Set("count", c.Int("count")+1)
            }),
    ).WithClass("p-8")
}
```

## Core Concepts

### Pages

Pages are functions that return UI:

```go
func MyPage(c *ctx.Context) ui.UI {
    return ui.Div(...)
}
```

### State

State is managed through Context:

```go
// Read
name := c.String("name")
count := c.Int("count")
active := c.Bool("active")

// Write
c.Set("name", "John")
c.Set("count", 42)

// Persist across sessions
c.Persist("user_id")
```

### Events

Attach handlers to elements:

```go
ui.Button(ui.T("Click")).
    WithID("my-btn").
    OnClick(c, func(c *ctx.Context) {
        c.Set("clicked", true)
    })

ui.Input().
    WithID("my-input").
    OnInput(c, func(c *ctx.Context) {
        c.Set("value", c.InputValue())
    })
```

### Routing

```go
app := server.New()

// Static routes
app.Route("/", HomePage)
app.Route("/about", AboutPage)

// Dynamic routes
app.Route("/user/:id", UserPage)
app.Route("/post/:slug", PostPage)

// Access params
func UserPage(c *ctx.Context) ui.UI {
    userID := c.Params["id"]
    return ui.Div(ui.T("User: " + userID))
}
```

### Layouts

```go
app.Layout("/", func(c *ctx.Context, child ui.UI) ui.UI {
    return ui.Div(
        ui.Nav(...),
        child,
        ui.Footer(...),
    )
})

app.Layout("/admin", AdminLayout) // applies to /admin/*
```

## UI Elements

### Basic Elements

```go
ui.Div(children...)
ui.Span(children...)
ui.P(children...)
ui.H1(children...) // H1-H4
ui.Button(children...)
ui.A(children...)
ui.Ul(children...)
ui.Li(children...)
ui.Form(children...)
ui.Input()
ui.Img()
ui.T("text") // text node
```

### Attributes

```go
ui.Div().
    WithID("my-id").
    WithClass("my-class").
    WithStyle("color: red").
    WithAttr("data-value", "123")
```

### Components

```go
// Layout
ui.Stack("16px", children...)  // vertical
ui.Row("8px", children...)     // horizontal
ui.Card(children...)

// Feedback
ui.Alert("Message", "success") // success|info|warning|error
ui.Badge("Label")
ui.Spinner()
ui.Progress(75)

// Data
ui.Table(headers, rows)

// Interactive
ui.Modal("id", c, children...)
ui.Tabs("id", c, items)
ui.Dropdown("id", c, trigger, items)
```

## Styling

### Tailwind CSS

```go
ui.EnableTailwind()

ui.Div().WithClass("flex items-center gap-4 p-6 bg-white rounded-lg shadow")
```

### Inline Styles

```go
ui.Div().WithS(ui.S().
    Flex().
    Gap("16px").
    P("20px").
    Bg("#fff").
    Rounded("8px"))
```

### Custom CSS

```go
ui.AddCSS(`
.my-component {
    background: linear-gradient(to right, #667eea, #764ba2);
}
`)
```

### Animations

```go
ui.Card(...).Animate("fadeIn")   // fadeIn, slideUp, scaleIn
ui.Card(...).Animate("bounce")   // bounce, pulse, spin
ui.Card(...).Hover("lift")       // lift, scale, glow
```

## Development

```go
// Development server with hot reload
app := server.NewDev(".")

// Production server
app := server.New()
```

## Architecture

```
Browser                          Server
┌─────────────────┐             ┌─────────────────┐
│  WASM Client    │◄──────────►│  Go Server      │
│  (424KB)        │  WebSocket  │                 │
│                 │             │  ┌───────────┐  │
│  - DOM Patching │             │  │  Context  │  │
│  - Event Fwd    │             │  │  (State)  │  │
└─────────────────┘             │  └───────────┘  │
                                │        ↓        │
                                │  ┌───────────┐  │
                                │  │  Render   │  │
                                │  │  + Diff   │  │
                                │  └───────────┘  │
                                └─────────────────┘
```

## License

MIT
