# Breadcrumbs

Navigation breadcrumbs for hierarchical pages.

## Usage

```go
ui.Breadcrumbs([]ui.BreadcrumbItem{
    {Label: "Home", Href: "/"},
    {Label: "Products", Href: "/products"},
    {Label: "Electronics", Href: "/products/electronics"},
    {Label: "Laptops"},  // Current page (no href)
})
```

## Output

```
Home / Products / Electronics / Laptops
```

- Links are clickable except the last item
- Last item shows as plain text (current page)

## Example

```go
func ProductPage(c *forge.Context) ui.UI {
    product := getProduct(c.Param("id"))
    category := getCategory(product.CategoryID)
    
    return ui.Div(
        ui.Breadcrumbs([]ui.BreadcrumbItem{
            {Label: "Home", Href: "/"},
            {Label: category.Name, Href: "/category/" + category.ID},
            {Label: product.Name},
        }),
        ui.H1(ui.T(product.Name)),
        // ...
    )
}
```

## Styling

Add breadcrumb styles:

```go
ui.Style(ui.BreadcrumbStyles)
```

Or customize:

```css
.breadcrumbs { display: flex; gap: 8px; font-size: 14px; }
.breadcrumb-link { color: #3b82f6; }
.breadcrumb-sep { color: #9ca3af; }
.breadcrumb-current { color: #6b7280; }
```
