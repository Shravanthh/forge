package ui

import "github.com/Shravanthh/forge/ctx"

// VirtualList renders only visible items for large lists.
// height: container height, itemHeight: each item height, items: total items
func VirtualList(id string, c *ctx.Context, height, itemHeight int, items []UI) Element {
	scrollTop := c.Int("vl_" + id + "_scroll")
	
	visibleCount := height/itemHeight + 2
	startIdx := scrollTop / itemHeight
	if startIdx < 0 {
		startIdx = 0
	}
	endIdx := startIdx + visibleCount
	if endIdx > len(items) {
		endIdx = len(items)
	}

	var visible []UI
	for i := startIdx; i < endIdx; i++ {
		visible = append(visible, Div(items[i]).WithStyle(
			"position:absolute;top:"+itoa(i*itemHeight)+"px;height:"+itoa(itemHeight)+"px;width:100%",
		))
	}

	return Div(
		Div(visible...).WithStyle("position:relative;height:" + itoa(len(items)*itemHeight) + "px"),
	).WithID(id).
		WithStyle("height:"+itoa(height)+"px;overflow-y:auto;position:relative").
		WithAttr("data-forge-scroll", id).
		OnScroll(c, func(c *ctx.Context) {
			c.Set("vl_"+id+"_scroll", c.Int("_scroll_top"))
		})
}

// OnScroll attaches a scroll event handler.
func (e Element) OnScroll(c *ctx.Context, h ctx.EventHandler) Element {
	return e.withEvent(c, "scroll", h)
}
