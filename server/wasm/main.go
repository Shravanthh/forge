//go:build js && wasm

package main

import (
	"encoding/json"
	"syscall/js"
)

var ws js.Value
var sessionID string

type Patch struct {
	Type  string            `json:"type"`
	ID    string            `json:"id"`
	HTML  string            `json:"html,omitempty"`
	Attrs map[string]string `json:"attrs,omitempty"`
	Text  string            `json:"text,omitempty"`
}

type Message struct {
	Type    string  `json:"type"`
	ID      string  `json:"id,omitempty"`
	Patches []Patch `json:"patches,omitempty"`
}

func main() {
	doc := js.Global().Get("document")

	// Wait for DOM ready
	if doc.Get("readyState").String() == "loading" {
		done := make(chan struct{})
		var cb js.Func
		cb = js.FuncOf(func(this js.Value, args []js.Value) any {
			cb.Release()
			close(done)
			return nil
		})
		doc.Call("addEventListener", "DOMContentLoaded", cb)
		<-done
	}

	// Setup event delegation
	setupEvents(doc)

	// Connect WebSocket
	connect()

	// Keep alive
	select {}
}

func setupEvents(doc js.Value) {
	// Click events
	doc.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		evt := args[0]
		target := evt.Get("target")
		
		el := closest(target, "[data-forge-click]")
		if !el.IsNull() {
			if el.Get("type").String() != "checkbox" {
				evt.Call("preventDefault")
			}
			send(el.Get("dataset").Get("forgeClick").String(), "")
		}
		return nil
	}))

	// Input events
	doc.Call("addEventListener", "input", js.FuncOf(func(this js.Value, args []js.Value) any {
		target := args[0].Get("target")
		id := target.Get("dataset").Get("forgeInput")
		if !id.IsUndefined() {
			send(id.String(), target.Get("value").String())
		}
		return nil
	}))

	// Keydown events
	doc.Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) any {
		evt := args[0]
		if evt.Get("key").String() == "Enter" {
			target := evt.Get("target")
			id := target.Get("dataset").Get("forgeKeydown")
			if !id.IsUndefined() {
				evt.Call("preventDefault")
				send(id.String(), target.Get("value").String())
			}
		}
		return nil
	}))

	// Change events
	doc.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) any {
		target := args[0].Get("target")
		id := target.Get("dataset").Get("forgeChange")
		if !id.IsUndefined() {
			val := target.Get("value").String()
			if target.Get("type").String() == "checkbox" {
				val = "false"
				if target.Get("checked").Bool() {
					val = "true"
				}
			}
			send(id.String(), val)
		}
		return nil
	}))

	// Scroll events (debounced)
	var scrollTimer js.Value
	doc.Call("addEventListener", "scroll", js.FuncOf(func(this js.Value, args []js.Value) any {
		target := args[0].Get("target")
		id := target.Get("dataset").Get("forgeScroll")
		if !id.IsUndefined() {
			if !scrollTimer.IsUndefined() {
				js.Global().Call("clearTimeout", scrollTimer)
			}
			scrollTimer = js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) any {
				sendScroll(id.String(), target.Get("scrollTop").Int())
				return nil
			}), 100)
		}
		return nil
	}), true)

	// Drag and drop
	var dragID string
	doc.Call("addEventListener", "dragstart", js.FuncOf(func(this js.Value, args []js.Value) any {
		target := args[0].Get("target")
		id := target.Get("dataset").Get("forgeDrag")
		if !id.IsUndefined() {
			dragID = id.String()
		}
		return nil
	}))

	doc.Call("addEventListener", "dragover", js.FuncOf(func(this js.Value, args []js.Value) any {
		evt := args[0]
		target := evt.Get("target")
		if closest(target, "[data-forge-dropzone]").Truthy() {
			evt.Call("preventDefault")
		}
		return nil
	}))

	doc.Call("addEventListener", "drop", js.FuncOf(func(this js.Value, args []js.Value) any {
		evt := args[0]
		evt.Call("preventDefault")
		target := closest(evt.Get("target"), "[data-forge-dropzone]")
		if !target.IsNull() {
			id := target.Get("dataset").Get("forgeDrop")
			if !id.IsUndefined() && dragID != "" {
				sendDrop(id.String(), dragID)
				dragID = ""
			}
		}
		return nil
	}))
}

func closest(el js.Value, selector string) js.Value {
	if el.IsNull() || el.IsUndefined() {
		return js.Null()
	}
	return el.Call("closest", selector)
}

func connect() {
	loc := js.Global().Get("location")
	proto := "ws:"
	if loc.Get("protocol").String() == "https:" {
		proto = "wss:"
	}
	host := loc.Get("host").String()
	url := proto + "//" + host + "/ws"
	if sessionID != "" {
		url += "?session=" + sessionID
	}

	ws = js.Global().Get("WebSocket").New(url)

	ws.Set("onmessage", js.FuncOf(func(this js.Value, args []js.Value) any {
		data := args[0].Get("data").String()
		var msg Message
		if err := json.Unmarshal([]byte(data), &msg); err != nil {
			return nil
		}

		if msg.Type == "session" {
			sessionID = msg.ID
		} else if msg.Type == "patch" {
			for _, p := range msg.Patches {
				applyPatch(p)
			}
		} else if msg.Type == "reload" {
			js.Global().Get("location").Call("reload")
		}
		return nil
	}))

	ws.Set("onclose", js.FuncOf(func(this js.Value, args []js.Value) any {
		// Reconnect after 1 second
		js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, args []js.Value) any {
			connect()
			return nil
		}), 1000)
		return nil
	}))
}

func send(id, value string) {
	if ws.Get("readyState").Int() != 1 {
		return
	}
	msg := map[string]string{"type": "event", "id": id, "value": value}
	data, _ := json.Marshal(msg)
	ws.Call("send", string(data))
}

func sendScroll(id string, scrollTop int) {
	if ws.Get("readyState").Int() != 1 {
		return
	}
	msg := map[string]any{"type": "scroll", "id": id, "scrollTop": scrollTop}
	data, _ := json.Marshal(msg)
	ws.Call("send", string(data))
}

func sendDrop(id, dragID string) {
	if ws.Get("readyState").Int() != 1 {
		return
	}
	msg := map[string]string{"type": "drop", "id": id, "dragId": dragID}
	data, _ := json.Marshal(msg)
	ws.Call("send", string(data))
}

func applyPatch(p Patch) {
	doc := js.Global().Get("document")
	el := doc.Call("querySelector", "[data-forge-id=\""+p.ID+"\"]")

	switch p.Type {
	case "replace":
		if !el.IsNull() {
			morph(el, p.HTML)
		}
	case "attrs":
		if !el.IsNull() {
			for k, v := range p.Attrs {
				if k == "checked" {
					el.Set("checked", v != "")
				} else if v == "" {
					el.Call("removeAttribute", k)
				} else {
					el.Call("setAttribute", k, v)
				}
			}
		}
	case "text":
		// Text nodes don't have data-forge-id, find parent element
		if el.IsNull() {
			parentID := parentPath(p.ID)
			el = doc.Call("querySelector", "[data-forge-id=\""+parentID+"\"]")
		}
		if !el.IsNull() {
			el.Set("textContent", p.Text)
		}
	case "remove":
		if !el.IsNull() {
			el.Call("remove")
		}
	case "insert":
		// Parse parent ID from path
		parentID := parentPath(p.ID)
		parent := doc.Call("querySelector", "[data-forge-id=\""+parentID+"\"]")
		if !parent.IsNull() {
			tpl := doc.Call("createElement", "template")
			tpl.Set("innerHTML", p.HTML)
			node := tpl.Get("content").Get("firstChild")
			parent.Call("appendChild", node)
		}
	}
}

func morph(target js.Value, html string) {
	doc := js.Global().Get("document")
	tpl := doc.Call("createElement", "template")
	tpl.Set("innerHTML", html)
	src := tpl.Get("content").Get("firstChild")

	if src.IsNull() {
		return
	}

	// Different tag - replace entirely
	if target.Get("tagName").String() != src.Get("tagName").String() {
		target.Call("replaceWith", src)
		return
	}

	// Sync attributes
	syncAttrs(target, src)

	// Sync children
	morphChildren(target, src)
}

func syncAttrs(target, src js.Value) {
	srcAttrs := src.Get("attributes")
	targetAttrs := target.Get("attributes")

	// Build set of src attr names
	srcNames := make(map[string]bool)
	for i := 0; i < srcAttrs.Length(); i++ {
		name := srcAttrs.Index(i).Get("name").String()
		srcNames[name] = true
	}

	// Remove attrs not in src
	toRemove := []string{}
	for i := 0; i < targetAttrs.Length(); i++ {
		name := targetAttrs.Index(i).Get("name").String()
		if !srcNames[name] {
			toRemove = append(toRemove, name)
		}
	}
	for _, name := range toRemove {
		target.Call("removeAttribute", name)
	}

	// Set attrs from src
	for i := 0; i < srcAttrs.Length(); i++ {
		attr := srcAttrs.Index(i)
		name := attr.Get("name").String()
		value := attr.Get("value").String()
		if target.Call("getAttribute", name).String() != value {
			target.Call("setAttribute", name, value)
		}
	}

	// Handle input properties
	if target.Get("tagName").String() == "INPUT" {
		target.Set("checked", src.Call("hasAttribute", "checked").Bool())
		if src.Call("hasAttribute", "value").Bool() {
			target.Set("value", src.Call("getAttribute", "value").String())
		}
	}
}

func morphChildren(target, src js.Value) {
	tKids := target.Get("childNodes")
	sKids := src.Get("childNodes")
	tLen := tKids.Length()
	sLen := sKids.Length()

	max := tLen
	if sLen > max {
		max = sLen
	}

	for i := 0; i < max; i++ {
		var t, s js.Value
		if i < tLen {
			t = tKids.Index(i)
		}
		if i < sLen {
			s = sKids.Index(i)
		}

		if s.IsUndefined() || s.IsNull() {
			if !t.IsUndefined() && !t.IsNull() {
				t.Call("remove")
			}
			continue
		}

		if t.IsUndefined() || t.IsNull() {
			target.Call("appendChild", s.Call("cloneNode", true))
			continue
		}

		tType := t.Get("nodeType").Int()
		sType := s.Get("nodeType").Int()

		if tType != sType {
			t.Call("replaceWith", s.Call("cloneNode", true))
			continue
		}

		if tType == 3 { // Text node
			if t.Get("textContent").String() != s.Get("textContent").String() {
				t.Set("textContent", s.Get("textContent").String())
			}
			continue
		}

		if tType == 1 { // Element
			morph(t, s.Get("outerHTML").String())
		}
	}
}

func parentPath(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return "0"
}
