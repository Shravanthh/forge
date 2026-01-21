package server

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/Shravanthh/forge/ctx"
	"github.com/Shravanthh/forge/render"
	"github.com/Shravanthh/forge/ui"
)

//go:embed wasm/forge.wasm
var forgeWASM []byte

//go:embed wasm/wasm_exec.js
var wasmExecJS []byte

// App is the main Forge application.
type App struct {
	sessions *SessionManager
	router   *Router
	uploads  map[string]UploadHandler
}

// LayoutFunc wraps a page with layout.
type LayoutFunc func(*ctx.Context, ui.UI) ui.UI

// New creates a new Forge application.
func New() *App {
	return &App{
		sessions: NewSessionManager(nil),
		router:   NewRouter(),
		uploads:  make(map[string]UploadHandler),
	}
}

// Route registers a page handler.
func (a *App) Route(path string, page PageFunc) { a.router.Add(path, page) }

// Layout registers a layout for a path prefix.
func (a *App) Layout(prefix string, layout LayoutFunc) { a.router.AddLayout(prefix, layout) }

// ServeHTTP implements http.Handler.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// File uploads
	if handler, ok := a.uploads[path]; ok {
		handleUpload(w, r, handler)
		return
	}

	switch path {
	case "/forge.wasm":
		w.Header().Set("Content-Type", "application/wasm")
		w.Write(forgeWASM)
		return
	case "/ws":
		pagePath := r.URL.Query().Get("path")
		if pagePath == "" {
			pagePath = "/"
		}
		page, params := a.router.Match(pagePath)
		if page == nil {
			page, params = a.router.Match("/")
		}
		a.sessions.HandleWebSocket(page, params)(w, r)
		return
	}

	page, params := a.router.Match(path)
	if page == nil {
		http.NotFound(w, r)
		return
	}

	c := ctx.New()
	c.Params = params
	ui.ResetEventCounter()
	content := page(c)

	for _, layout := range a.router.GetLayouts(path) {
		content = layout(c, content)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, wrapHTML(render.HTML(content)))
}

func wrapHTML(body string) string {
	var headScripts, bodyScripts strings.Builder
	for _, s := range ui.GetHeadScripts() {
		headScripts.WriteString(s)
		headScripts.WriteByte('\n')
	}
	for _, s := range ui.GetBodyScripts() {
		bodyScripts.WriteString(s)
		bodyScripts.WriteByte('\n')
	}

	return `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>Forge App</title>
<style>` + ui.GetCSS() + `</style>
` + headScripts.String() + `
</head>
<body>
` + body + `
` + bodyScripts.String() + `
<script>` + string(wasmExecJS) + `
const go=new Go();WebAssembly.instantiateStreaming(fetch("/forge.wasm"),go.importObject).then(r=>go.run(r.instance));
</script>
</body>
</html>`
}

// Run starts the server.
func (a *App) Run(addr string) error {
	fmt.Printf("Forge running at http://localhost%s\n", addr)
	return http.ListenAndServe(addr, a)
}
