# Lazy Loading

Defer rendering of components until they're visible.

## Lazy Components

Load heavy components only when scrolled into view:

```go
func Page(c *forge.Context) ui.UI {
    return ui.Div(
        ui.H1(ui.T("Dashboard")),
        
        // Loads immediately
        ui.Card("Stats", StatsWidget(c)),
        
        // Loads when scrolled into view
        ui.Lazy("chart", c, func() ui.UI {
            return ExpensiveChart(c)
        }),
        
        ui.Lazy("table", c, func() ui.UI {
            return LargeDataTable(c)
        }),
    )
}
```

## How It Works

1. Initially renders a placeholder with shimmer animation
2. Uses IntersectionObserver to detect visibility
3. When visible, triggers server to render actual content
4. Replaces placeholder with real component

## Lazy Images

Load images only when visible:

```go
ui.LazyImage("/images/large-photo.jpg", "Photo description")
```

- Shows transparent placeholder initially
- Loads actual image when scrolled into view
- Fades in smoothly when loaded

## OnVisible Event

Trigger any action when element becomes visible:

```go
ui.Div(
    ui.T("Analytics section"),
).OnVisible(c, func(c *forge.Context) {
    // Track view, load data, etc.
    trackPageSection("analytics")
})
```

## Styling

Add lazy loading styles:

```go
ui.Style(ui.LazyStyles)
```

Includes:
- Shimmer animation for placeholders
- Fade-in transition for images
- Minimum height to prevent layout shift

## Example: Infinite Scroll

```go
func Feed(c *forge.Context) ui.UI {
    posts := getPosts(c.Int("page"), 10)
    
    var items []ui.UI
    for _, post := range posts {
        items = append(items, PostCard(post))
    }
    
    // Load more trigger
    items = append(items, ui.Div().
        WithID("load-more").
        OnVisible(c, func(c *forge.Context) {
            c.Set("page", c.Int("page")+1)
        }))
    
    return ui.Div(items...)
}
```
