# Pagination

Paginate lists and tables.

## Usage

```go
func UsersPage(c *forge.Context) ui.UI {
    perPage := 10
    total := getTotalUsers()
    page := ui.GetPage(c, "users")
    offset := ui.GetOffset(c, "users", perPage)
    
    users := getUsers(offset, perPage)
    
    return ui.Div(
        ui.Table(/* render users */),
        ui.Pagination("users", c, total, perPage),
    )
}
```

## API

### Pagination Component

```go
ui.Pagination(id string, c *ctx.Context, total int, perPage int) Element
```

- `id` - Unique identifier for this pagination
- `c` - Context
- `total` - Total number of items
- `perPage` - Items per page

### Helper Functions

```go
// Get current page (1-indexed)
page := ui.GetPage(c, "users") // 1, 2, 3...

// Get offset for database queries
offset := ui.GetOffset(c, "users", 10) // 0, 10, 20...
```

## Example with Database

```go
func ProductsPage(c *forge.Context) ui.UI {
    perPage := 20
    
    // Get total count
    var total int
    db.QueryRow("SELECT COUNT(*) FROM products").Scan(&total)
    
    // Get current page items
    offset := ui.GetOffset(c, "products", perPage)
    rows, _ := db.Query(
        "SELECT * FROM products LIMIT ? OFFSET ?",
        perPage, offset,
    )
    
    products := scanProducts(rows)
    
    return ui.Div(
        renderProductList(products),
        ui.Pagination("products", c, total, perPage),
    )
}
```

## Styling

Add pagination styles:

```go
ui.Style(ui.PaginationStyles)
```

Or customize:

```css
.pagination { display: flex; gap: 4px; }
.page-btn { min-width: 36px; height: 36px; border-radius: 6px; }
.page-btn.active { background: #3b82f6; color: white; }
.page-btn.disabled { opacity: 0.5; }
```
