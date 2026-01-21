package ui

// BreadcrumbItem represents a breadcrumb link.
type BreadcrumbItem struct {
	Label string
	Href  string
}

// Breadcrumbs creates a breadcrumb navigation.
func Breadcrumbs(items []BreadcrumbItem) Element {
	var children []UI

	for i, item := range items {
		if i > 0 {
			children = append(children, Span(T("/")).WithClass("breadcrumb-sep"))
		}

		if i == len(items)-1 {
			// Last item (current page)
			children = append(children, Span(T(item.Label)).WithClass("breadcrumb-current"))
		} else {
			children = append(children, A(T(item.Label)).WithAttr("href", item.Href).WithClass("breadcrumb-link"))
		}
	}

	return Nav(children...).WithClass("breadcrumbs")
}

// BreadcrumbStyles contains CSS for breadcrumbs.
const BreadcrumbStyles = `
.breadcrumbs{display:flex;align-items:center;gap:8px;font-size:14px}
.breadcrumb-link{color:#3b82f6;text-decoration:none}
.breadcrumb-link:hover{text-decoration:underline}
.breadcrumb-sep{color:#9ca3af}
.breadcrumb-current{color:#6b7280}
`
