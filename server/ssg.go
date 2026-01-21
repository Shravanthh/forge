package server

import (
	"os"
	"path/filepath"

	"github.com/Shravanthh/forge/ctx"
	"github.com/Shravanthh/forge/render"
	"github.com/Shravanthh/forge/ui"
)

// StaticPage defines a page to be statically generated.
type StaticPage struct {
	Path   string
	Params map[string]string
}

// GenerateStatic generates static HTML files for the given pages.
// Output files are written to the outDir directory.
func (a *App) GenerateStatic(outDir string, pages []StaticPage) error {
	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	for _, sp := range pages {
		page, _ := a.router.Match(sp.Path)
		if page == nil {
			continue
		}

		c := ctx.New()
		c.Params = sp.Params
		ui.ResetEventCounter()
		content := page(c)

		for _, layout := range a.router.GetLayouts(sp.Path) {
			content = layout(c, content)
		}

		html := wrapHTMLStatic(render.HTML(content))

		outPath := filepath.Join(outDir, sp.Path)
		if sp.Path == "/" {
			outPath = filepath.Join(outDir, "index.html")
		} else {
			outPath = filepath.Join(outDir, sp.Path, "index.html")
		}

		os.MkdirAll(filepath.Dir(outPath), 0755)
		if err := os.WriteFile(outPath, []byte(html), 0644); err != nil {
			return err
		}
	}
	return nil
}

func wrapHTMLStatic(body string) string {
	return `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>Forge App</title>
<style>` + ui.GetCSS() + `</style>
</head>
<body>
` + body + `
</body>
</html>`
}
