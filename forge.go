// Package forge provides a Go-native UI framework for building web applications.
//
// Example:
//
//	package main
//
//	import (
//	    "forge"
//	    "forge/ui"
//	)
//
//	func main() {
//	    app := forge.New()
//	    app.Route("/", func(c *forge.Context) ui.UI {
//	        return ui.Div(ui.T("Hello, Forge!"))
//	    })
//	    app.Run(":3000")
//	}
package forge

import (
	"forge/ctx"
	"forge/server"
)

// Context is the request context containing state and params.
type Context = ctx.Context

// App is the Forge application.
type App = server.App

// DevServer is the development server with hot reload.
type DevServer = server.DevServer

// PageFunc is a function that renders a page.
type PageFunc = server.PageFunc

// LayoutFunc wraps pages with a layout.
type LayoutFunc = server.LayoutFunc

// New creates a new Forge application.
func New() *App { return server.New() }

// NewDev creates a development server with hot reload.
func NewDev(watchDir string) *DevServer { return server.NewDev(watchDir) }
