# Routing

Forge supports static routes, dynamic parameters, and layouts.

## Basic Routing

```go
app := forge.New()

app.Route("/", HomePage)
app.Route("/about", AboutPage)
app.Route("/contact", ContactPage)

app.Run(":3000")
```

## Dynamic Routes

Use `:param` for dynamic segments:

```go
app.Route("/user/:id", UserPage)
app.Route("/post/:slug", PostPage)
app.Route("/category/:cat/item/:id", ItemPage)

func UserPage(c *forge.Context) ui.UI {
    userID := c.Params["id"]
    return ui.H1(ui.T("User: " + userID))
}

func ItemPage(c *forge.Context) ui.UI {
    category := c.Params["cat"]
    itemID := c.Params["id"]
    // ...
}
```

## Layouts

Layouts wrap pages with common UI (header, footer, sidebar):

```go
app.Layout("/", RootLayout)           // applies to all routes
app.Layout("/admin", AdminLayout)     // applies to /admin/*
app.Layout("/dashboard", DashLayout)  // applies to /dashboard/*

func RootLayout(c *forge.Context, child ui.UI) ui.UI {
    return ui.Div(
        ui.Header(
            ui.Nav(
                ui.A(ui.T("Home")).WithAttr("href", "/"),
                ui.A(ui.T("About")).WithAttr("href", "/about"),
            ),
        ).WithClass("navbar"),
        ui.Main(child).WithClass("content"),
        ui.Footer(ui.T("© 2026")).WithClass("footer"),
    )
}

func AdminLayout(c *forge.Context, child ui.UI) ui.UI {
    return ui.Div(
        ui.Div(ui.T("Admin Sidebar")).WithClass("sidebar"),
        ui.Div(child).WithClass("admin-content"),
    ).WithClass("admin-layout")
}
```

## Navigation

### Links

```go
ui.A(ui.T("Go to About")).WithAttr("href", "/about")
```

### Programmatic Navigation

Currently, use standard links. Client-side navigation without page reload is handled automatically by the WASM client.

## 404 Handling

Unmatched routes return a 404 response. Add a catch-all route if needed:

```go
// Add last
app.Route("/:path", NotFoundPage)

func NotFoundPage(c *forge.Context) ui.UI {
    return ui.Div(
        ui.H1(ui.T("404 - Not Found")),
        ui.A(ui.T("Go Home")).WithAttr("href", "/"),
    )
}
```

## Route Organization

For larger apps, organize routes in separate files:

```
my-app/
├── main.go
├── routes/
│   ├── home.go
│   ├── user.go
│   └── admin.go
```

```go
// routes/user.go
package routes

func RegisterUserRoutes(app *forge.App) {
    app.Route("/user/:id", UserPage)
    app.Route("/user/:id/edit", EditUserPage)
}

// main.go
func main() {
    app := forge.New()
    routes.RegisterUserRoutes(app)
    app.Run(":3000")
}
```
