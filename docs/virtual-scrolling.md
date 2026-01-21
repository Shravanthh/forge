# Virtual Scrolling

Efficiently render large lists with virtual scrolling.

## Basic Usage

```go
func LargeList(c *forge.Context) ui.UI {
    // Generate 10,000 items
    var items []ui.UI
    for i := 0; i < 10000; i++ {
        items = append(items, ui.Div(
            ui.T(fmt.Sprintf("Item %d", i)),
        ).WithClass("list-item"))
    }
    
    // Only renders visible items
    return ui.VirtualList(
        "my-list",     // unique ID
        c,             // context
        400,           // container height (px)
        50,            // item height (px)
        items,         // all items
    )
}
```

## How It Works

1. Container has fixed height with `overflow-y: auto`
2. Inner container has full height (items × itemHeight)
3. Only visible items are rendered with absolute positioning
4. Scroll events update which items are visible

## Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | string | Unique identifier for the list |
| `c` | *Context | Forge context |
| `height` | int | Container height in pixels |
| `itemHeight` | int | Height of each item in pixels |
| `items` | []UI | All items to render |

## Styling

```go
ui.AddCSS(`
.list-item {
    padding: 12px 16px;
    border-bottom: 1px solid #eee;
    display: flex;
    align-items: center;
}

.list-item:hover {
    background: #f5f5f5;
}
`)
```

## With Dynamic Content

```go
func UserList(c *forge.Context) ui.UI {
    users := getUsers() // fetch from database
    
    var items []ui.UI
    for _, user := range users {
        items = append(items, UserRow(user))
    }
    
    return ui.VirtualList("users", c, 600, 72, items)
}

func UserRow(user User) ui.UI {
    return ui.Row("12px",
        ui.Avatar(user.Avatar, user.Name),
        ui.Stack("4px",
            ui.Span(ui.T(user.Name)).WithClass("font-bold"),
            ui.Span(ui.T(user.Email)).WithClass("text-gray-500 text-sm"),
        ),
    ).WithClass("list-item")
}
```

## Limitations

- All items must have the same height
- Horizontal scrolling not supported
- Search/filter requires re-rendering the list

## Performance Tips

1. Keep item components simple
2. Avoid complex calculations in item render
3. Use fixed heights (no auto-sizing)
4. Consider pagination for extremely large datasets (100k+)
