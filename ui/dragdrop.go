package ui

import "github.com/Shravanthh/forge/ctx"

// Draggable makes an element draggable.
func (e Element) Draggable(id string) Element {
	return e.WithAttr("draggable", "true").WithAttr("data-forge-drag", id)
}

// DropZone makes an element a drop target.
func (e Element) DropZone(c *ctx.Context, onDrop func(c *ctx.Context, dragID string)) Element {
	handlerID := "drop_" + e.ID
	c.On(handlerID, func(c *ctx.Context) {
		dragID := c.String("_drag_id")
		if dragID != "" && onDrop != nil {
			onDrop(c, dragID)
		}
	})
	if e.Events == nil {
		e.Events = make(map[string]string)
	}
	e.Events["drop"] = handlerID
	return e.WithAttr("data-forge-dropzone", "true")
}

// SortableList creates a drag-and-drop sortable list.
func SortableList(id string, c *ctx.Context, items []UI, onReorder func(c *ctx.Context, fromIdx, toIdx int)) Element {
	var children []UI
	for i, item := range items {
		idx := i
		child := Div(item).
			WithID(id + "_" + itoa(i)).
			Draggable(itoa(i)).
			DropZone(c, func(c *ctx.Context, dragID string) {
				fromIdx := atoi(dragID)
				if fromIdx != idx && onReorder != nil {
					onReorder(c, fromIdx, idx)
				}
			})
		children = append(children, child)
	}
	return Div(children...).WithID(id).WithClass("sortable-list")
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}
