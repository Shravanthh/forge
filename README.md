# Forge

A Go-native UI framework for building modern web applications.

[![Go Reference](https://pkg.go.dev/badge/github.com/Shravanthh/forge.svg)](https://pkg.go.dev/github.com/Shravanthh/forge)

## Features

- **Pure Go** - Write UI entirely in Go, no JavaScript required
- **Server-Driven** - State lives on server, UI updates via WebSocket
- **Live Updates** - Real-time DOM patching without page reloads
- **Tailwind CSS** - Built-in Tailwind support
- **Type Safe** - Full Go type safety for your UI code
- **Fast** - 425KB WASM client, sub-millisecond updates

## Installation

```bash
go install github.com/Shravanthh/forge/cmd/forge@latest
```

## Quick Start

```bash
forge new my-app
cd my-app
forge dev
```

Open http://localhost:3000

## Documentation

- [Getting Started](docs/getting-started.md)
- [UI Elements](docs/elements.md)
- [State Management](docs/state.md)
- [Routing](docs/routing.md)
- [Styling](docs/styling.md)
- [Components](docs/components.md)
- [Events](docs/events.md)
- [Forms](docs/forms.md)
- [File Uploads](docs/uploads.md)
- [Drag & Drop](docs/dragdrop.md)
- [Virtual Scrolling](docs/virtual-scrolling.md)
- [Static Site Generation](docs/ssg.md)
- [Third-Party Integration](docs/third-party.md)
- [Deployment](docs/deployment.md)
- [API Reference](docs/api.md)

## Example

```go
package main

import (
    "github.com/Shravanthh/forge"
    "github.com/Shravanthh/forge/ui"
)

func main() {
    app := forge.New()
    app.Route("/", HomePage)
    app.Run(":3000")
}

func HomePage(c *forge.Context) ui.UI {
    count := c.Int("count")
    
    return ui.Div(
        ui.H1(ui.T("Counter")),
        ui.P(ui.T(fmt.Sprintf("Count: %d", count))),
        ui.Button(ui.T("Increment")).
            WithID("inc").
            WithClass("btn btn-primary").
            OnClick(c, func(c *forge.Context) {
                c.Set("count", c.Int("count")+1)
            }),
    ).WithClass("container")
}
```

## Architecture

```
Browser                          Server
┌─────────────────┐             ┌─────────────────┐
│  WASM Client    │◄──────────►│  Go Server      │
│  (425KB)        │  WebSocket  │                 │
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

## CLI Commands

| Command | Description |
|---------|-------------|
| `forge new <name>` | Create new project |
| `forge dev` | Start dev server with hot reload |
| `forge start` | Start production server |
| `forge build` | Build for production |
| `forge test` | Run tests |

## License

MIT
