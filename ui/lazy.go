package ui

import "github.com/Shravanthh/forge/ctx"

// Lazy defers rendering until the component is visible.
func Lazy(id string, c *ctx.Context, loader func() UI) Element {
	loaded := c.Bool(id + "_loaded")

	if loaded {
		return Div(loader()).WithID(id).WithClass("lazy-loaded")
	}

	// Placeholder with intersection observer trigger
	return Div(
		Div().WithClass("lazy-placeholder"),
	).WithID(id).
		WithClass("lazy-container").
		OnVisible(c, func(c *ctx.Context) {
			c.Set(id+"_loaded", true)
		})
}

// OnVisible registers a handler when element becomes visible.
func (e Element) OnVisible(c *ctx.Context, handler ctx.EventHandler) Element {
	return e.withEvent(c, "visible", handler)
}

// LazyImage loads image when visible.
func LazyImage(src, alt string) Element {
	return Img().
		WithAttr("data-src", src).
		WithAttr("alt", alt).
		WithClass("lazy-image").
		WithAttr("src", "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
}

// LazyStyles contains CSS for lazy loading.
const LazyStyles = `
.lazy-container{min-height:100px}
.lazy-placeholder{background:linear-gradient(90deg,#f0f0f0 25%,#e0e0e0 50%,#f0f0f0 75%);background-size:200% 100%;animation:shimmer 1.5s infinite;height:100%;min-height:100px;border-radius:8px}
@keyframes shimmer{0%{background-position:200% 0}100%{background-position:-200% 0}}
.lazy-image{opacity:0;transition:opacity 0.3s}
.lazy-image.loaded{opacity:1}
`
