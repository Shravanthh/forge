// Package diff computes the differences between two UI trees
// and generates patches for efficient DOM updates.
package diff

import (
	"strconv"

	"github.com/Shravanthh/forge/render"
	"github.com/Shravanthh/forge/ui"
)

// PatchType defines the type of DOM modification.
type PatchType string

const (
	Replace    PatchType = "replace" // Replace entire element
	UpdateAttr PatchType = "attrs"   // Update attributes only
	UpdateText PatchType = "text"    // Update text content
	Insert     PatchType = "insert"  // Insert new element
	Remove     PatchType = "remove"  // Remove element
)

// Patch represents a single DOM modification to be applied by the client.
type Patch struct {
	Type  PatchType         `json:"type"`            // Type of patch
	ID    string            `json:"id"`              // Target element ID
	HTML  string            `json:"html,omitempty"`  // New HTML (for replace/insert)
	Attrs map[string]string `json:"attrs,omitempty"` // Changed attributes
	Text  string            `json:"text,omitempty"`  // New text content
}

// Diff compares two UI trees and returns the minimal set of patches
// needed to transform the old tree into the new tree.
//
//	oldUI := ui.Div(ui.T("Hello"))
//	newUI := ui.Div(ui.T("World"))
//	patches := diff.Diff(oldUI, newUI)
//	// [{Type: "text", ID: "0.0", Text: "World"}]
func Diff(oldUI, newUI ui.UI) []Patch {
	return diffNode(oldUI, newUI, "0")
}

func diffNode(oldN, newN ui.UI, path string) []Patch {
	if oldN == nil && newN == nil {
		return nil
	}
	if oldN == nil {
		return []Patch{{Type: Insert, ID: path, HTML: render.HTML(newN)}}
	}
	if newN == nil {
		id := path
		if e, ok := oldN.(ui.Element); ok && e.ID != "" {
			id = e.ID
		}
		return []Patch{{Type: Remove, ID: id}}
	}
	if nodeType(oldN) != nodeType(newN) {
		return []Patch{{Type: Replace, ID: path, HTML: render.HTML(newN)}}
	}

	switch o := oldN.(type) {
	case ui.Text:
		if n := newN.(ui.Text); o.Value != n.Value {
			return []Patch{{Type: UpdateText, ID: path, Text: n.Value}}
		}
	case ui.Element:
		return diffElement(o, newN.(ui.Element), path)
	case ui.Raw:
		if n := newN.(ui.Raw); o.HTML != n.HTML {
			return []Patch{{Type: Replace, ID: path, HTML: n.HTML}}
		}
	}
	return nil
}

func diffElement(old, new ui.Element, path string) []Patch {
	id := path
	if new.ID != "" {
		id = new.ID
	}

	if old.Tag != new.Tag {
		return []Patch{{Type: Replace, ID: id, HTML: render.HTML(new)}}
	}

	var patches []Patch
	if attrs := diffAttrs(old, new); len(attrs) > 0 {
		patches = append(patches, Patch{Type: UpdateAttr, ID: id, Attrs: attrs})
	}
	patches = append(patches, diffChildren(old.Children, new.Children, path)...)
	return patches
}

func diffAttrs(old, new ui.Element) map[string]string {
	changes := make(map[string]string)
	if old.Class != new.Class {
		changes["class"] = new.Class
	}
	if old.Style != new.Style {
		changes["style"] = new.Style
	}
	for k, v := range new.Attrs {
		if old.Attrs[k] != v {
			changes[k] = v
		}
	}
	for k := range old.Attrs {
		if _, ok := new.Attrs[k]; !ok {
			changes[k] = ""
		}
	}
	for k, v := range new.Events {
		if old.Events[k] != v {
			changes["data-forge-"+k] = v
		}
	}
	for k := range old.Events {
		if _, ok := new.Events[k]; !ok {
			changes["data-forge-"+k] = ""
		}
	}
	return changes
}

func diffChildren(oldC, newC []ui.UI, parentPath string) []Patch {
	var patches []Patch

	oldByID := make(map[string]ui.UI)
	for _, c := range oldC {
		if e, ok := c.(ui.Element); ok && e.ID != "" {
			oldByID[e.ID] = c
		}
	}

	usedIDs := make(map[string]bool)
	for i, newChild := range newC {
		childPath := parentPath + "." + strconv.Itoa(i)
		var oldChild ui.UI

		if e, ok := newChild.(ui.Element); ok && e.ID != "" {
			if o, found := oldByID[e.ID]; found {
				oldChild = o
				usedIDs[e.ID] = true
			}
		}
		if oldChild == nil && i < len(oldC) {
			o := oldC[i]
			if e, ok := o.(ui.Element); ok && e.ID != "" {
				if !usedIDs[e.ID] {
					oldChild = nil
				}
			} else {
				oldChild = o
			}
		}
		patches = append(patches, diffNode(oldChild, newChild, childPath)...)
	}

	for _, oldChild := range oldC {
		if e, ok := oldChild.(ui.Element); ok && e.ID != "" && !usedIDs[e.ID] {
			patches = append(patches, Patch{Type: Remove, ID: e.ID})
		}
	}
	return patches
}

func nodeType(n ui.UI) int {
	switch n.(type) {
	case ui.Element:
		return 1
	case ui.Text:
		return 2
	case ui.Raw:
		return 3
	}
	return 0
}
