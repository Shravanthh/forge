package ui

var headScripts []string
var bodyScripts []string

// AddHeadScript adds a script to the <head> section.
// Use for third-party libraries that need early loading.
func AddHeadScript(script string) {
	headScripts = append(headScripts, script)
}

// AddBodyScript adds a script to the end of <body>.
// Use for widgets and scripts that need DOM ready.
func AddBodyScript(script string) {
	bodyScripts = append(bodyScripts, script)
}

// GetHeadScripts returns all registered head scripts.
func GetHeadScripts() []string { return headScripts }

// GetBodyScripts returns all registered body scripts.
func GetBodyScripts() []string { return bodyScripts }

// ResetScripts clears all registered scripts.
func ResetScripts() {
	headScripts = nil
	bodyScripts = nil
}

// Embed creates a container for third-party widget.
// The widget JS should target this element by ID.
func Embed(id string) Element {
	return Div().WithID(id).WithClass("forge-embed")
}

// IFrame creates an iframe element for embedding external content.
func IFrame(src string) Element {
	return El("iframe").
		WithAttr("src", src).
		WithAttr("frameborder", "0").
		WithStyle("border:none")
}
