package ui

func el(tag string, children ...UI) Element { return Element{Tag: tag, Children: children} }

// Layout elements

// Div creates a <div> element.
func Div(children ...UI) Element { return el("div", children...) }

// Span creates a <span> element.
func Span(children ...UI) Element { return el("span", children...) }

// Main creates a <main> element.
func Main(children ...UI) Element { return el("main", children...) }

// Section creates a <section> element.
func Section(children ...UI) Element { return el("section", children...) }

// Article creates an <article> element.
func Article(children ...UI) Element { return el("article", children...) }

// Header creates a <header> element.
func Header(children ...UI) Element { return el("header", children...) }

// Footer creates a <footer> element.
func Footer(children ...UI) Element { return el("footer", children...) }

// Nav creates a <nav> element.
func Nav(children ...UI) Element { return el("nav", children...) }

// Text elements

// P creates a <p> paragraph element.
func P(children ...UI) Element { return el("p", children...) }

// H1 creates an <h1> heading element.
func H1(children ...UI) Element { return el("h1", children...) }

// H2 creates an <h2> heading element.
func H2(children ...UI) Element { return el("h2", children...) }

// H3 creates an <h3> heading element.
func H3(children ...UI) Element { return el("h3", children...) }

// H4 creates an <h4> heading element.
func H4(children ...UI) Element { return el("h4", children...) }

// Interactive elements

// Button creates a <button> element.
func Button(children ...UI) Element { return el("button", children...) }

// A creates an <a> anchor element.
func A(children ...UI) Element { return el("a", children...) }

// Form creates a <form> element.
func Form(children ...UI) Element { return el("form", children...) }

// Label creates a <label> element.
func Label(children ...UI) Element { return el("label", children...) }

// List elements

// Ul creates an <ul> unordered list element.
func Ul(children ...UI) Element { return el("ul", children...) }

// Li creates an <li> list item element.
func Li(children ...UI) Element { return el("li", children...) }

// Table elements

// Table creates a <table> element.
func Table(children ...UI) Element { return el("table", children...) }

// Thead creates a <thead> element.
func Thead(children ...UI) Element { return el("thead", children...) }

// Tbody creates a <tbody> element.
func Tbody(children ...UI) Element { return el("tbody", children...) }

// Tr creates a <tr> table row element.
func Tr(children ...UI) Element { return el("tr", children...) }

// Th creates a <th> table header cell element.
func Th(children ...UI) Element { return el("th", children...) }

// Td creates a <td> table data cell element.
func Td(children ...UI) Element { return el("td", children...) }

// Form elements

// Select creates a <select> element.
func Select(children ...UI) Element { return el("select", children...) }

// Option creates an <option> element.
func Option(children ...UI) Element { return el("option", children...) }

// Textarea creates a <textarea> element.
func Textarea() Element { return Element{Tag: "textarea"} }

// Self-closing elements

// Input creates an <input> element.
func Input() Element { return Element{Tag: "input"} }

// Img creates an <img> element.
func Img() Element { return Element{Tag: "img"} }

// Br creates a <br> line break element.
func Br() Element { return Element{Tag: "br"} }

// Hr creates an <hr> horizontal rule element.
func Hr() Element { return Element{Tag: "hr"} }

// El creates a custom element with the specified tag.
func El(tag string, children ...UI) Element { return el(tag, children...) }

// T creates a text node with the given string value.
//
//	ui.P(ui.T("Hello, World!"))
func T(s string) Text { return Text{Value: s} }

// Element builder methods

// WithID sets the element's ID attribute.
func (e Element) WithID(id string) Element { e.ID = id; return e }

// WithClass sets the element's class attribute.
func (e Element) WithClass(c string) Element { e.Class = c; return e }

// WithStyle sets the element's inline style attribute.
func (e Element) WithStyle(s string) Element { e.Style = s; return e }

// WithS sets the element's style using a Style builder.
func (e Element) WithS(s Style) Element { e.Style = s.String(); return e }

// WithAttr sets a custom attribute on the element.
func (e Element) WithAttr(k, v string) Element {
	if e.Attrs == nil {
		e.Attrs = make(map[string]string)
	}
	e.Attrs[k] = v
	return e
}

// WithChildren appends children to the element.
func (e Element) WithChildren(children ...UI) Element {
	e.Children = append(e.Children, children...)
	return e
}
