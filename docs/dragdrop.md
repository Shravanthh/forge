# Drag & Drop

Implement drag and drop in Forge.

## Draggable Elements

Make any element draggable:

```go
ui.Div(
    ui.T("Drag me"),
).WithID("item-1").Draggable("item-1")
```

## Drop Zones

Create drop targets:

```go
ui.Div(
    ui.T("Drop here"),
).WithID("drop-zone").DropZone(c, func(c *forge.Context, dragID string) {
    // dragID is the ID passed to Draggable()
    fmt.Println("Dropped:", dragID)
})
```

## Sortable List

Create a reorderable list:

```go
func TaskList(c *forge.Context) ui.UI {
    tasks := c.Get("tasks").([]string)
    
    var items []ui.UI
    for _, task := range tasks {
        items = append(items, ui.Div(ui.T(task)).WithClass("task-item"))
    }
    
    return ui.SortableList("tasks", c, items, func(c *forge.Context, from, to int) {
        tasks := c.Get("tasks").([]string)
        
        // Reorder
        item := tasks[from]
        tasks = append(tasks[:from], tasks[from+1:]...)
        tasks = append(tasks[:to], append([]string{item}, tasks[to:]...)...)
        
        c.Set("tasks", tasks)
    })
}
```

## Kanban Board Example

```go
func KanbanBoard(c *forge.Context) ui.UI {
    columns := []string{"todo", "doing", "done"}
    
    var cols []ui.UI
    for _, col := range columns {
        col := col
        tasks := getTasksForColumn(c, col)
        
        var taskItems []ui.UI
        for _, task := range tasks {
            task := task
            taskItems = append(taskItems,
                ui.Div(ui.T(task.Title)).
                    WithID("task-"+task.ID).
                    Draggable(task.ID).
                    WithClass("task-card"),
            )
        }
        
        cols = append(cols, ui.Div(
            ui.H3(ui.T(strings.Title(col))),
            ui.Div(taskItems...).
                WithID("col-"+col).
                DropZone(c, func(c *forge.Context, taskID string) {
                    moveTask(c, taskID, col)
                }).
                WithClass("column"),
        ).WithClass("kanban-column"))
    }
    
    return ui.Row("16px", cols...).WithClass("kanban-board")
}
```

## Styling Drag & Drop

```go
ui.AddCSS(`
.sortable-list {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

[draggable="true"] {
    cursor: grab;
}

[draggable="true"]:active {
    cursor: grabbing;
}

[data-forge-dropzone="true"] {
    min-height: 100px;
    border: 2px dashed #ccc;
    border-radius: 8px;
    transition: border-color 0.2s;
}

[data-forge-dropzone="true"]:hover {
    border-color: #3b82f6;
}
`)
```

## Limitations

- Drag & drop involves server roundtrips, so there may be slight latency
- For complex drag interactions (like drawing), consider a hybrid approach
- Visual feedback during drag is CSS-only (no JS drag preview)
