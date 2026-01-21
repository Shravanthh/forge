# Getting Started

## Prerequisites

- Go 1.25 or later
- A terminal

## Installation

Install the Forge CLI:

```bash
go install github.com/Shravanthh/forge/cmd/forge@latest
```

Verify installation:

```bash
forge version
```

## Create Your First App

```bash
forge new my-app
cd my-app
forge dev
```

Open http://localhost:3000 in your browser.

## Project Structure

```
my-app/
├── main.go         # Application entry point
├── go.mod          # Go module file
├── pages/          # Page components (optional)
└── components/     # Reusable components (optional)
```

## Basic Example

```go
package main

import (
    "github.com/Shravanthh/forge"
    "github.com/Shravanthh/forge/ui"
)

func init() {
    ui.EnableTailwind()
}

func main() {
    app := forge.New()
    app.Route("/", HomePage)
    app.Run(":3000")
}

func HomePage(c *forge.Context) ui.UI {
    return ui.Div(
        ui.H1(ui.T("Hello, Forge!")),
        ui.P(ui.T("Welcome to your first Forge app.")),
    ).WithClass("p-8")
}
```

## Development vs Production

### Development (with hot reload)

```go
app := forge.NewDev(".")
```

Or use CLI:

```bash
forge dev
```

### Production

```go
app := forge.New()
```

Or use CLI:

```bash
forge build
./my-app
```

## Next Steps

- [UI Elements](elements.md) - Learn about available elements
- [State Management](state.md) - Handle application state
- [Routing](routing.md) - Add multiple pages
- [Styling](styling.md) - Style your app
