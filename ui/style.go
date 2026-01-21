package ui

import "strings"

// Style is a chainable CSS style builder.
type Style map[string]string

func (s Style) String() string {
	var b strings.Builder
	for k, v := range s {
		if b.Len() > 0 {
			b.WriteByte(';')
		}
		b.WriteString(k)
		b.WriteByte(':')
		b.WriteString(v)
	}
	return b.String()
}

// S creates a new Style builder.
func S() Style { return make(Style) }

func (s Style) set(k, v string) Style { s[k] = v; return s }

// Layout
func (s Style) Display(v string) Style        { return s.set("display", v) }
func (s Style) Flex() Style                   { return s.Display("flex") }
func (s Style) Grid() Style                   { return s.Display("grid") }
func (s Style) Block() Style                  { return s.Display("block") }
func (s Style) None() Style                   { return s.Display("none") }
func (s Style) FlexDir(v string) Style        { return s.set("flex-direction", v) }
func (s Style) FlexRow() Style                { return s.FlexDir("row") }
func (s Style) FlexCol() Style                { return s.FlexDir("column") }
func (s Style) JustifyContent(v string) Style { return s.set("justify-content", v) }
func (s Style) AlignItems(v string) Style     { return s.set("align-items", v) }
func (s Style) Gap(v string) Style            { return s.set("gap", v) }

// Spacing
func (s Style) P(v string) Style  { return s.set("padding", v) }
func (s Style) Px(v string) Style { return s.set("padding-left", v).set("padding-right", v) }
func (s Style) Py(v string) Style { return s.set("padding-top", v).set("padding-bottom", v) }
func (s Style) M(v string) Style  { return s.set("margin", v) }
func (s Style) Mx(v string) Style { return s.set("margin-left", v).set("margin-right", v) }
func (s Style) My(v string) Style { return s.set("margin-top", v).set("margin-bottom", v) }

// Size
func (s Style) W(v string) Style    { return s.set("width", v) }
func (s Style) H(v string) Style    { return s.set("height", v) }
func (s Style) MaxW(v string) Style { return s.set("max-width", v) }
func (s Style) MinH(v string) Style { return s.set("min-height", v) }

// Colors
func (s Style) Bg(v string) Style    { return s.set("background", v) }
func (s Style) Color(v string) Style { return s.set("color", v) }

// Typography
func (s Style) Font(v string) Style       { return s.set("font-family", v) }
func (s Style) FontSize(v string) Style   { return s.set("font-size", v) }
func (s Style) FontWeight(v string) Style { return s.set("font-weight", v) }
func (s Style) Bold() Style               { return s.FontWeight("bold") }
func (s Style) TextAlign(v string) Style  { return s.set("text-align", v) }
func (s Style) Center() Style             { return s.TextAlign("center") }
func (s Style) LineThrough() Style        { return s.set("text-decoration", "line-through") }

// Border
func (s Style) Border(v string) Style       { return s.set("border", v) }
func (s Style) BorderRadius(v string) Style { return s.set("border-radius", v) }
func (s Style) Rounded(v string) Style      { return s.BorderRadius(v) }

// Effects
func (s Style) Shadow(v string) Style     { return s.set("box-shadow", v) }
func (s Style) Opacity(v string) Style    { return s.set("opacity", v) }
func (s Style) Transition(v string) Style { return s.set("transition", v) }

// Position
func (s Style) Pos(v string) Style { return s.set("position", v) }
func (s Style) Absolute() Style    { return s.Pos("absolute") }
func (s Style) Relative() Style    { return s.Pos("relative") }
func (s Style) Fixed() Style       { return s.Pos("fixed") }

// Cursor
func (s Style) Cursor(v string) Style { return s.set("cursor", v) }
func (s Style) Pointer() Style        { return s.Cursor("pointer") }
