package ui

import (
	"github.com/Shravanthh/forge/ctx"
)

// Pagination creates a pagination component.
func Pagination(id string, c *ctx.Context, total, perPage int) Element {
	currentPage := c.Int(id + "_page")
	if currentPage < 1 {
		currentPage = 1
	}

	totalPages := (total + perPage - 1) / perPage
	if totalPages < 1 {
		totalPages = 1
	}

	var buttons []UI

	// Previous
	prevBtn := Button(T("←")).WithID(id + "-prev").WithClass("page-btn")
	if currentPage > 1 {
		prevBtn = prevBtn.OnClick(c, func(c *ctx.Context) {
			c.Set(id+"_page", c.Int(id+"_page")-1)
		})
	} else {
		prevBtn = prevBtn.WithAttr("disabled", "true").WithClass("page-btn disabled")
	}
	buttons = append(buttons, prevBtn)

	// Page numbers
	start := currentPage - 2
	if start < 1 {
		start = 1
	}
	end := start + 4
	if end > totalPages {
		end = totalPages
		start = end - 4
		if start < 1 {
			start = 1
		}
	}

	for i := start; i <= end; i++ {
		page := i
		cls := "page-btn"
		if page == currentPage {
			cls += " active"
		}
		buttons = append(buttons, Button(T(itoa(page))).
			WithID(id+"-page-"+itoa(page)).
			WithClass(cls).
			OnClick(c, func(c *ctx.Context) {
				c.Set(id+"_page", page)
			}))
	}

	// Next
	nextBtn := Button(T("→")).WithID(id + "-next").WithClass("page-btn")
	if currentPage < totalPages {
		nextBtn = nextBtn.OnClick(c, func(c *ctx.Context) {
			c.Set(id+"_page", c.Int(id+"_page")+1)
		})
	} else {
		nextBtn = nextBtn.WithAttr("disabled", "true").WithClass("page-btn disabled")
	}
	buttons = append(buttons, nextBtn)

	return Div(buttons...).WithClass("pagination")
}

// GetPage returns current page (1-indexed).
func GetPage(c *ctx.Context, id string) int {
	page := c.Int(id + "_page")
	if page < 1 {
		return 1
	}
	return page
}

// GetOffset returns offset for database queries.
func GetOffset(c *ctx.Context, id string, perPage int) int {
	return (GetPage(c, id) - 1) * perPage
}

// PaginationStyles contains CSS for pagination.
const PaginationStyles = `
.pagination{display:flex;gap:4px;align-items:center}
.page-btn{min-width:36px;height:36px;border:1px solid #e5e7eb;background:#fff;border-radius:6px;cursor:pointer;font-size:14px}
.page-btn:hover:not(.disabled){background:#f3f4f6}
.page-btn.active{background:#3b82f6;color:#fff;border-color:#3b82f6}
.page-btn.disabled{opacity:0.5;cursor:not-allowed}
`
