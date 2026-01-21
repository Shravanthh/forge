package ui

import "github.com/Shravanthh/forge/ctx"

import "strconv"

func itoa(i int) string { return strconv.Itoa(i) }

// Layout components

// Stack creates a vertical flex container.
func Stack(gap string, children ...UI) Element {
	return Div(children...).WithS(S().Flex().FlexCol().Gap(gap))
}

// Row creates a horizontal flex container.
func Row(gap string, children ...UI) Element {
	return Div(children...).WithS(S().Flex().FlexRow().Gap(gap))
}

// Card creates a card container.
func Card(children ...UI) Element {
	return Div(children...).WithClass("card")
}

// Divider creates a horizontal line.
func Divider() Element { return Hr().WithClass("divider") }

// Feedback components

// Alert creates an alert message.
func Alert(msg, variant string) Element {
	return Div(T(msg)).WithClass("alert alert-" + variant)
}

// Badge creates a small label.
func Badge(text string) Element {
	return Span(T(text)).WithClass("badge")
}

// Spinner creates a loading spinner.
func Spinner() Element { return Div().WithClass("spinner") }

// Progress creates a progress bar (0-100).
func Progress(value int) Element {
	return Div(
		Div().WithClass("progress-bar").WithStyle("width:" + itoa(value) + "%"),
	).WithClass("progress")
}

// Avatar creates an avatar image.
func Avatar(src, alt string) Element {
	return Img().WithAttr("src", src).WithAttr("alt", alt).WithClass("avatar")
}

// Data components

// DataTable creates a table from headers and rows.
func DataTable(headers []string, rows [][]UI) Element {
	var ths []UI
	for _, h := range headers {
		ths = append(ths, Th(T(h)))
	}
	thead := Thead(Tr(ths...))

	var trs []UI
	for _, row := range rows {
		var tds []UI
		for _, cell := range row {
			tds = append(tds, Td(cell))
		}
		trs = append(trs, Tr(tds...))
	}
	tbody := Tbody(trs...)

	return Table(thead, tbody).WithClass("table")
}

// Interactive components

// Modal creates a modal dialog.
func Modal(id string, c *ctx.Context, children ...UI) Element {
	if !c.Bool("modal_" + id) {
		return Div().WithID("modal-" + id).WithStyle("display:none")
	}
	return Div(
		Div(
			Div(children...).WithClass("modal-content"),
		).WithClass("modal-backdrop").WithID("modal-"+id).OnClick(c, func(c *ctx.Context) {
			c.Set("modal_"+id, false)
		}),
	).WithClass("modal")
}

// OpenModal returns a handler to open a modal.
func OpenModal(id string) func(*ctx.Context) {
	return func(c *ctx.Context) { c.Set("modal_"+id, true) }
}

// CloseModal returns a handler to close a modal.
func CloseModal(id string) func(*ctx.Context) {
	return func(c *ctx.Context) { c.Set("modal_"+id, false) }
}

// TabItem represents a tab.
type TabItem struct {
	Key     string
	Label   string
	Content UI
}

// Tabs creates a tabbed interface.
func Tabs(id string, c *ctx.Context, tabs []TabItem) Element {
	active := c.String("tab_" + id)
	if active == "" && len(tabs) > 0 {
		active = tabs[0].Key
	}

	var btns []UI
	var content UI
	for _, tab := range tabs {
		cls := "tab-btn"
		if tab.Key == active {
			cls += " active"
			content = tab.Content
		}
		key := tab.Key
		btns = append(btns, Button(T(tab.Label)).WithClass(cls).WithID("tab-"+id+"-"+key).OnClick(c, func(c *ctx.Context) {
			c.Set("tab_"+id, key)
		}))
	}

	return Div(
		Div(btns...).WithClass("tab-list"),
		Div(content).WithClass("tab-content"),
	).WithClass("tabs")
}

// DropdownItem represents a dropdown menu item.
type DropdownItem struct {
	Key     string
	Label   string
	OnClick func(*ctx.Context)
}

// Dropdown creates a dropdown menu.
func Dropdown(id string, c *ctx.Context, trigger UI, items []DropdownItem) Element {
	isOpen := c.Bool("dropdown_" + id)

	triggerEl := Div(trigger).WithClass("dropdown-trigger").WithID("dd-"+id).OnClick(c, func(c *ctx.Context) {
		c.Set("dropdown_"+id, !c.Bool("dropdown_"+id))
	})

	var menu Element
	if isOpen {
		var menuItems []UI
		for _, item := range items {
			item := item
			menuItems = append(menuItems, Div(T(item.Label)).WithClass("dropdown-item").WithID("dd-"+id+"-"+item.Key).OnClick(c, func(c *ctx.Context) {
				c.Set("dropdown_"+id, false)
				if item.OnClick != nil {
					item.OnClick(c)
				}
			}))
		}
		menu = Div(menuItems...).WithClass("dropdown-menu")
	} else {
		menu = Div().WithClass("dropdown-menu").WithStyle("display:none")
	}

	return Div(triggerEl, menu).WithClass("dropdown")
}
