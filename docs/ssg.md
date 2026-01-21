# Static Site Generation (SSG)

Generate static HTML files for deployment without a server.

## Basic Usage

```go
func main() {
    app := forge.New()
    app.Route("/", HomePage)
    app.Route("/about", AboutPage)
    app.Route("/contact", ContactPage)
    
    // Generate static files
    err := app.GenerateStatic("./dist", []server.StaticPage{
        {Path: "/"},
        {Path: "/about"},
        {Path: "/contact"},
    })
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Static site generated in ./dist")
}
```

## Output Structure

```
dist/
├── index.html          # /
├── about/
│   └── index.html      # /about
└── contact/
    └── index.html      # /contact
```

## Dynamic Routes

For dynamic routes, specify all possible paths:

```go
// Get all blog posts
posts := getAllPosts()

var pages []server.StaticPage
pages = append(pages, server.StaticPage{Path: "/"})
pages = append(pages, server.StaticPage{Path: "/blog"})

for _, post := range posts {
    pages = append(pages, server.StaticPage{
        Path:   "/blog/" + post.Slug,
        Params: map[string]string{"slug": post.Slug},
    })
}

app.GenerateStatic("./dist", pages)
```

## CLI Command

Add to your CLI or build script:

```go
// cmd/generate/main.go
func main() {
    app := setupApp()
    pages := getPages()
    
    if err := app.GenerateStatic("./dist", pages); err != nil {
        log.Fatal(err)
    }
}
```

Run:

```bash
go run cmd/generate/main.go
```

## Deployment

### Netlify

```toml
# netlify.toml
[build]
  command = "go run cmd/generate/main.go"
  publish = "dist"
```

### Vercel

```json
{
  "buildCommand": "go run cmd/generate/main.go",
  "outputDirectory": "dist"
}
```

### GitHub Pages

```yaml
# .github/workflows/deploy.yml
name: Deploy
on:
  push:
    branches: [main]

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: '1.25'
      - run: go run cmd/generate/main.go
      - uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./dist
```

## Limitations

- No WebSocket support (static HTML only)
- No server-side state
- No event handlers (buttons won't work)
- Best for content-only pages

## Hybrid Approach

Use SSG for content pages, server for interactive pages:

```go
// Static pages (content only)
app.GenerateStatic("./dist", []server.StaticPage{
    {Path: "/"},
    {Path: "/about"},
    {Path: "/blog"},
})

// Serve interactive pages from server
app.Route("/app", DashboardPage)  // needs interactivity
app.Route("/login", LoginPage)    // needs forms
```
