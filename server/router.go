package server

import "strings"

// Route holds route info.
type Route struct {
	Pattern  string
	Page     PageFunc
	segments []string
	params   []string
}

// Router handles URL routing.
type Router struct {
	routes  []*Route
	layouts map[string]LayoutFunc
}

// NewRouter creates a router.
func NewRouter() *Router {
	return &Router{layouts: make(map[string]LayoutFunc)}
}

// Add registers a route.
func (r *Router) Add(pattern string, page PageFunc) {
	route := &Route{Pattern: pattern, Page: page}
	route.segments, route.params = parsePattern(pattern)
	r.routes = append(r.routes, route)
}

// AddLayout registers a layout for a path prefix.
func (r *Router) AddLayout(prefix string, layout LayoutFunc) {
	r.layouts[prefix] = layout
}

// Match finds a matching route and extracts params.
func (r *Router) Match(path string) (PageFunc, map[string]string) {
	segs := splitPath(path)
	for _, route := range r.routes {
		if params, ok := matchRoute(route, segs); ok {
			return route.Page, params
		}
	}
	return nil, nil
}

// GetLayouts returns all matching layouts for a path.
func (r *Router) GetLayouts(path string) []LayoutFunc {
	var layouts []LayoutFunc
	for prefix, layout := range r.layouts {
		if strings.HasPrefix(path, prefix) {
			layouts = append(layouts, layout)
		}
	}
	return layouts
}

func parsePattern(pattern string) ([]string, []string) {
	segs := splitPath(pattern)
	params := make([]string, len(segs))
	for i, seg := range segs {
		if len(seg) > 0 && seg[0] == ':' {
			params[i] = seg[1:]
			segs[i] = ""
		}
	}
	return segs, params
}

func splitPath(path string) []string {
	path = strings.Trim(path, "/")
	if path == "" {
		return nil
	}
	return strings.Split(path, "/")
}

func matchRoute(route *Route, segs []string) (map[string]string, bool) {
	if len(route.segments) != len(segs) {
		if len(route.segments) == 0 && len(segs) == 0 {
			return map[string]string{}, true
		}
		return nil, false
	}
	params := make(map[string]string)
	for i, seg := range route.segments {
		if route.params[i] != "" {
			params[route.params[i]] = segs[i]
		} else if seg != segs[i] {
			return nil, false
		}
	}
	return params, true
}
