package render

import (
	"html"
	"strconv"
	"strings"

	"forge/ui"
)

var selfClosing = map[string]bool{
	"input": true, "img": true, "br": true, "hr": true,
	"meta": true, "link": true, "area": true, "base": true,
}

// HTML renders a UI tree to HTML string with forge IDs.
func HTML(node ui.UI) string {
	var b strings.Builder
	b.Grow(4096)
	renderNode(&b, node, "0")
	return b.String()
}

func renderNode(b *strings.Builder, node ui.UI, path string) {
	switch n := node.(type) {
	case ui.Element:
		renderElement(b, n, path)
	case ui.Text:
		b.WriteString(html.EscapeString(n.Value))
	case ui.Raw:
		b.WriteString(n.HTML)
	}
}

func renderElement(b *strings.Builder, e ui.Element, path string) {
	id := path
	if e.ID != "" {
		id = e.ID
	}

	b.WriteByte('<')
	b.WriteString(e.Tag)
	b.WriteString(` data-forge-id="`)
	b.WriteString(id)
	b.WriteByte('"')

	if e.Class != "" {
		b.WriteString(` class="`)
		b.WriteString(html.EscapeString(e.Class))
		b.WriteByte('"')
	}
	if e.Style != "" {
		b.WriteString(` style="`)
		b.WriteString(html.EscapeString(e.Style))
		b.WriteByte('"')
	}
	for k, v := range e.Attrs {
		b.WriteByte(' ')
		b.WriteString(k)
		b.WriteString(`="`)
		b.WriteString(html.EscapeString(v))
		b.WriteByte('"')
	}
	for evt, handler := range e.Events {
		b.WriteString(` data-forge-`)
		b.WriteString(evt)
		b.WriteString(`="`)
		b.WriteString(handler)
		b.WriteByte('"')
	}

	if selfClosing[e.Tag] {
		b.WriteString(" />")
		return
	}

	b.WriteByte('>')
	for i, child := range e.Children {
		renderNode(b, child, path+"."+strconv.Itoa(i))
	}
	b.WriteString("</")
	b.WriteString(e.Tag)
	b.WriteByte('>')
}
