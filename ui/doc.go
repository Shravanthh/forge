// Package ui provides the declarative UI DSL for building user interfaces in Go.
//
// # Elements
//
// Create HTML elements using constructor functions:
//
//	ui.Div(children...)     // <div>
//	ui.Span(children...)    // <span>
//	ui.Button(children...)  // <button>
//	ui.Input()              // <input />
//	ui.T("text")            // text node
//
// # Attributes
//
// Chain methods to add attributes:
//
//	ui.Div(
//	    ui.T("Hello"),
//	).WithID("main").
//	  WithClass("container").
//	  WithStyle("color: red").
//	  WithAttr("data-value", "123")
//
// # Events
//
// Attach event handlers to elements:
//
//	ui.Button(ui.T("Click")).
//	    WithID("btn").
//	    OnClick(c, func(c *ctx.Context) {
//	        c.Set("clicked", true)
//	    })
//
// Available events: OnClick, OnInput, OnChange, OnSubmit, OnKeydown
//
// # Styling
//
// Use the Style builder for inline styles:
//
//	ui.Div().WithS(ui.S().Flex().Gap("16px").P("20px").Bg("#fff"))
//
// Or use Tailwind CSS classes:
//
//	ui.EnableTailwind()
//	ui.Div().WithClass("flex gap-4 p-5 bg-white")
//
// # Components
//
// Pre-built components for common UI patterns:
//
//	ui.Card(children...)                    // Card container
//	ui.Modal("id", c, children...)          // Modal dialog
//	ui.Table(headers, rows)                 // Data table
//	ui.Tabs("id", c, items)                 // Tabbed interface
//	ui.Alert("message", "success")          // Alert message
//	ui.Badge("label")                       // Badge/tag
//	ui.Spinner()                            // Loading spinner
//	ui.Progress(75)                         // Progress bar
//
// # Animations
//
// Add animations to elements:
//
//	ui.Card(...).Animate("fadeIn")   // Entry animation
//	ui.Card(...).Hover("lift")       // Hover effect
package ui
