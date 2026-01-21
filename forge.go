// Package forge provides a Go-native UI framework for building modern web applications.
//
// Forge enables developers to build interactive web UIs entirely in Go, with no JavaScript required.
// It uses a server-driven architecture where state lives on the server and UI updates are sent
// to the browser via WebSocket.
//
// # Quick Start
//
//	package main
//
//	import (
//	    "github.com/Shravanthh/forge"
//	    "github.com/Shravanthh/forge/ui"
//	)
//
//	func main() {
//	    app := forge.New()
//	    app.Route("/", HomePage)
//	    app.Run(":3000")
//	}
//
//	func HomePage(c *forge.Context) ui.UI {
//	    count := c.Int("count")
//	    return ui.Div(
//	        ui.H1(ui.T("Counter")),
//	        ui.P(ui.T(fmt.Sprintf("Count: %d", count))),
//	        ui.Button(ui.T("Increment")).
//	            WithID("inc").
//	            OnClick(c, func(c *forge.Context) {
//	                c.Set("count", c.Int("count")+1)
//	            }),
//	    )
//	}
//
// # Architecture
//
// Forge follows a server-driven UI architecture:
//
//  1. Server renders UI to HTML and sends to browser
//  2. Browser loads minimal WASM client (425KB)
//  3. User interactions are sent to server via WebSocket
//  4. Server updates state, re-renders, and sends DOM patches
//  5. WASM client applies patches to update the DOM
//
// # Features
//
//   - Pure Go UI DSL with type safety
//   - Server-side state management
//   - Real-time updates via WebSocket
//   - Built-in Tailwind CSS support
//   - Dynamic routing with parameters
//   - Reusable components (Modal, Table, Tabs, etc.)
//   - CSS animations
//   - Hot reload development server
//   - CLI for project scaffolding
package forge

import (
	"github.com/Shravanthh/forge/ctx"
	"github.com/Shravanthh/forge/server"
)

// Context holds the request state, event handlers, and route parameters.
// It is passed to every page function and event handler.
//
// Use Context to read and write state:
//
//	count := c.Int("count")    // Read
//	c.Set("count", count + 1)  // Write
//	c.Persist("user_id")       // Persist across sessions
type Context = ctx.Context

// App is the main Forge application that handles routing and serves HTTP requests.
type App = server.App

// DevServer extends App with hot reload capability for development.
// It watches for file changes and automatically refreshes connected browsers.
type DevServer = server.DevServer

// PageFunc is a function that renders a page given a Context.
// Every route handler must implement this signature.
//
//	func HomePage(c *forge.Context) ui.UI {
//	    return ui.Div(ui.T("Hello"))
//	}
type PageFunc = server.PageFunc

// LayoutFunc wraps page content with a layout.
// Layouts are applied based on URL prefix matching.
//
//	app.Layout("/", func(c *forge.Context, child ui.UI) ui.UI {
//	    return ui.Div(
//	        ui.Nav(ui.T("My App")),
//	        child,
//	    )
//	})
type LayoutFunc = server.LayoutFunc

// New creates a new Forge application for production use.
//
//	app := forge.New()
//	app.Route("/", HomePage)
//	app.Run(":3000")
func New() *App { return server.New() }

// NewDev creates a development server with hot reload.
// Pass the directory to watch for file changes.
//
//	app := forge.NewDev(".")
//	app.Route("/", HomePage)
//	app.Run(":3000")
func NewDev(watchDir string) *DevServer { return server.NewDev(watchDir) }
