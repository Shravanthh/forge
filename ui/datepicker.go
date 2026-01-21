package ui

import (
	"time"

	"github.com/Shravanthh/forge/ctx"
)

// DatePicker creates a date picker component.
func DatePicker(id string, c *ctx.Context) Element {
	isOpen := c.Bool(id + "_open")
	selected := c.String(id)

	// Parse current view month/year
	viewYear := c.Int(id + "_year")
	viewMonth := c.Int(id + "_month")
	if viewYear == 0 {
		now := time.Now()
		viewYear = now.Year()
		viewMonth = int(now.Month())
	}

	// Input field
	input := Input().
		WithID(id).
		WithAttr("type", "text").
		WithAttr("value", selected).
		WithAttr("placeholder", "YYYY-MM-DD").
		WithAttr("readonly", "true").
		WithClass("datepicker-input").
		OnClick(c, func(c *ctx.Context) {
			c.Set(id+"_open", !c.Bool(id+"_open"))
			if c.Int(id+"_year") == 0 {
				now := time.Now()
				c.Set(id+"_year", now.Year())
				c.Set(id+"_month", int(now.Month()))
			}
		})

	if !isOpen {
		return Div(input).WithClass("datepicker")
	}

	// Calendar
	calendar := buildCalendar(id, c, viewYear, viewMonth, selected)

	return Div(
		input,
		Div(calendar...).WithClass("datepicker-dropdown").Animate("fadeIn"),
	).WithClass("datepicker")
}

func buildCalendar(id string, c *ctx.Context, year, month int, selected string) []UI {
	var elements []UI

	// Header with month/year navigation
	monthName := time.Month(month).String()
	header := Div(
		Button(T("‹")).WithID(id+"-prev-month").WithClass("dp-nav").OnClick(c, func(c *ctx.Context) {
			m := c.Int(id + "_month")
			y := c.Int(id + "_year")
			m--
			if m < 1 {
				m = 12
				y--
			}
			c.Set(id+"_month", m)
			c.Set(id+"_year", y)
		}),
		Span(T(monthName+" "+itoa(year))).WithClass("dp-title"),
		Button(T("›")).WithID(id+"-next-month").WithClass("dp-nav").OnClick(c, func(c *ctx.Context) {
			m := c.Int(id + "_month")
			y := c.Int(id + "_year")
			m++
			if m > 12 {
				m = 1
				y++
			}
			c.Set(id+"_month", m)
			c.Set(id+"_year", y)
		}),
	).WithClass("dp-header")
	elements = append(elements, header)

	// Day names
	days := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	var dayHeaders []UI
	for _, d := range days {
		dayHeaders = append(dayHeaders, Span(T(d)).WithClass("dp-day-name"))
	}
	elements = append(elements, Div(dayHeaders...).WithClass("dp-days-header"))

	// Calendar grid
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	startOffset := int(firstDay.Weekday())
	daysInMonth := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()

	var dayCells []UI

	// Empty cells for offset
	for i := 0; i < startOffset; i++ {
		dayCells = append(dayCells, Span(T("")).WithClass("dp-day empty"))
	}

	// Day cells
	for day := 1; day <= daysInMonth; day++ {
		d := day
		dateStr := formatDate(year, month, d)
		cls := "dp-day"
		if dateStr == selected {
			cls += " selected"
		}
		dayCells = append(dayCells, Button(T(itoa(d))).
			WithID(id+"-day-"+itoa(d)).
			WithClass(cls).
			OnClick(c, func(c *ctx.Context) {
				c.Set(id, formatDate(year, month, d))
				c.Set(id+"_open", false)
			}))
	}

	elements = append(elements, Div(dayCells...).WithClass("dp-grid"))

	return elements
}

func formatDate(year, month, day int) string {
	return itoa(year) + "-" + pad(month) + "-" + pad(day)
}

func pad(n int) string {
	if n < 10 {
		return "0" + itoa(n)
	}
	return itoa(n)
}

// DatePickerStyles contains CSS for date picker.
const DatePickerStyles = `
.datepicker{position:relative;display:inline-block}
.datepicker-input{padding:8px 12px;border:1px solid #d1d5db;border-radius:6px;cursor:pointer;min-width:150px}
.datepicker-dropdown{position:absolute;top:100%;left:0;margin-top:4px;background:#fff;border:1px solid #e5e7eb;border-radius:8px;box-shadow:0 4px 12px rgba(0,0,0,0.15);padding:12px;z-index:100;width:280px}
.dp-header{display:flex;justify-content:space-between;align-items:center;margin-bottom:12px}
.dp-title{font-weight:600}
.dp-nav{background:none;border:none;font-size:18px;cursor:pointer;padding:4px 8px;border-radius:4px}
.dp-nav:hover{background:#f3f4f6}
.dp-days-header{display:grid;grid-template-columns:repeat(7,1fr);gap:4px;margin-bottom:8px}
.dp-day-name{text-align:center;font-size:12px;color:#6b7280;font-weight:500}
.dp-grid{display:grid;grid-template-columns:repeat(7,1fr);gap:4px}
.dp-day{width:32px;height:32px;border:none;background:none;border-radius:6px;cursor:pointer;font-size:14px}
.dp-day:hover:not(.empty){background:#f3f4f6}
.dp-day.selected{background:#3b82f6;color:#fff}
.dp-day.empty{cursor:default}
`
