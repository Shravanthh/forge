// Package ui provides the UI DSL for building user interfaces in Go.
//
// Basic usage:
//
//	ui.Div(
//	    ui.H1(ui.T("Hello")),
//	    ui.P(ui.T("Welcome to Forge")),
//	    ui.Button(ui.T("Click me")).OnClick(c, handler),
//	).WithClass("container")
package ui
