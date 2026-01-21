package ui

// UI is the interface implemented by all renderable nodes.
// Elements, text nodes, and raw HTML all implement this interface.
type UI interface{ isUI() }

// Element represents an HTML element with tag, attributes, and children.
//
//	el := ui.Div(
//	    ui.H1(ui.T("Title")),
//	    ui.P(ui.T("Content")),
//	).WithClass("container")
type Element struct {
	Tag      string            // HTML tag name (e.g., "div", "button")
	ID       string            // Element ID (used for data-forge-id)
	Class    string            // CSS class attribute
	Style    string            // Inline style attribute
	Attrs    map[string]string // Additional attributes
	Children []UI              // Child nodes
	Events   map[string]string // Event handler IDs
}

func (Element) isUI() {}

// Text represents a text node. Use T() to create text nodes.
//
//	ui.P(ui.T("Hello, World!"))
type Text struct{ Value string }

func (Text) isUI() {}

// Raw represents raw HTML that will be inserted without escaping.
// Use with caution to avoid XSS vulnerabilities.
//
//	ui.Raw{HTML: "<strong>Bold</strong>"}
type Raw struct{ HTML string }

func (Raw) isUI() {}
