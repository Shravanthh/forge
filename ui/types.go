package ui

// UI is the interface for all renderable nodes.
type UI interface{ isUI() }

// Element represents an HTML element.
type Element struct {
	Tag      string
	ID       string
	Class    string
	Style    string
	Attrs    map[string]string
	Children []UI
	Events   map[string]string
}

func (Element) isUI() {}

// Text represents a text node.
type Text struct{ Value string }

func (Text) isUI() {}

// Raw represents raw HTML (use with caution).
type Raw struct{ HTML string }

func (Raw) isUI() {}
