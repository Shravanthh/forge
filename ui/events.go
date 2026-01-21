package ui

import (
	"fmt"
	"sync/atomic"

	"github.com/Shravanthh/forge/ctx"
)

var eventCounter uint64

// ResetEventCounter resets counter (call at start of each render)
func ResetEventCounter() { atomic.StoreUint64(&eventCounter, 0) }

func nextEventID() string {
	return fmt.Sprintf("e%d", atomic.AddUint64(&eventCounter, 1))
}

func (e Element) withEvent(c *ctx.Context, evtType string, handler ctx.EventHandler) Element {
	id := e.ID
	if id == "" {
		id = nextEventID()
	}
	handlerID := id + "_" + evtType
	c.On(handlerID, handler)
	if e.Events == nil {
		e.Events = make(map[string]string)
	}
	e.Events[evtType] = handlerID
	return e
}

// Event handlers
func (e Element) OnClick(c *ctx.Context, h ctx.EventHandler) Element   { return e.withEvent(c, "click", h) }
func (e Element) OnInput(c *ctx.Context, h ctx.EventHandler) Element   { return e.withEvent(c, "input", h) }
func (e Element) OnChange(c *ctx.Context, h ctx.EventHandler) Element  { return e.withEvent(c, "change", h) }
func (e Element) OnSubmit(c *ctx.Context, h ctx.EventHandler) Element  { return e.withEvent(c, "submit", h) }
func (e Element) OnKeydown(c *ctx.Context, h ctx.EventHandler) Element { return e.withEvent(c, "keydown", h) }
