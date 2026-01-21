package ui

func el(tag string, children ...UI) Element { return Element{Tag: tag, Children: children} }

// Layout elements
func Div(children ...UI) Element     { return el("div", children...) }
func Span(children ...UI) Element    { return el("span", children...) }
func Main(children ...UI) Element    { return el("main", children...) }
func Section(children ...UI) Element { return el("section", children...) }
func Article(children ...UI) Element { return el("article", children...) }
func Header(children ...UI) Element  { return el("header", children...) }
func Footer(children ...UI) Element  { return el("footer", children...) }
func Nav(children ...UI) Element     { return el("nav", children...) }

// Text elements
func P(children ...UI) Element  { return el("p", children...) }
func H1(children ...UI) Element { return el("h1", children...) }
func H2(children ...UI) Element { return el("h2", children...) }
func H3(children ...UI) Element { return el("h3", children...) }
func H4(children ...UI) Element { return el("h4", children...) }

// Interactive elements
func Button(children ...UI) Element { return el("button", children...) }
func A(children ...UI) Element      { return el("a", children...) }
func Form(children ...UI) Element   { return el("form", children...) }
func Label(children ...UI) Element  { return el("label", children...) }

// List elements
func Ul(children ...UI) Element { return el("ul", children...) }
func Li(children ...UI) Element { return el("li", children...) }

// Table elements
func Tr(children ...UI) Element { return el("tr", children...) }
func Th(children ...UI) Element { return el("th", children...) }
func Td(children ...UI) Element { return el("td", children...) }

// Self-closing elements
func Input() Element { return Element{Tag: "input"} }
func Img() Element   { return Element{Tag: "img"} }
func Br() Element    { return Element{Tag: "br"} }
func Hr() Element    { return Element{Tag: "hr"} }

// El creates a custom element
func El(tag string, children ...UI) Element { return el(tag, children...) }

// T creates a text node
func T(s string) Text { return Text{Value: s} }

// Element builder methods
func (e Element) WithID(id string) Element    { e.ID = id; return e }
func (e Element) WithClass(c string) Element  { e.Class = c; return e }
func (e Element) WithStyle(s string) Element  { e.Style = s; return e }
func (e Element) WithS(s Style) Element       { e.Style = s.String(); return e }
func (e Element) WithAttr(k, v string) Element {
	if e.Attrs == nil {
		e.Attrs = make(map[string]string)
	}
	e.Attrs[k] = v
	return e
}
func (e Element) WithChildren(children ...UI) Element {
	e.Children = append(e.Children, children...)
	return e
}
