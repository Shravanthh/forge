# HTTP Client

Make API calls from event handlers.

## Basic Usage

### GET Request

```go
ui.Button(ui.T("Load Data")).OnClick(c, func(c *forge.Context) {
    var users []User
    err := c.FetchJSON("https://api.example.com/users", &users)
    if err != nil {
        c.Set("error", err.Error())
        return
    }
    c.Set("users", users)
})
```

### POST Request

```go
ui.Button(ui.T("Submit")).OnClick(c, func(c *forge.Context) {
    payload := map[string]string{
        "name":  c.String("name"),
        "email": c.String("email"),
    }
    
    var result Response
    err := c.PostJSON("https://api.example.com/users", payload, &result)
    if err != nil {
        c.Set("error", err.Error())
        return
    }
    c.Set("success", true)
})
```

## Context Methods

```go
// GET request, returns raw bytes
data, err := c.Fetch(url)

// GET request, unmarshal JSON
err := c.FetchJSON(url, &result)

// POST request with JSON body, returns raw bytes
data, err := c.Post(url, body)

// POST request, unmarshal JSON response
err := c.PostJSON(url, body, &result)

// Custom request
data, err := c.Request(method, url, body, headers)
```

## API Client

For repeated calls to the same API:

```go
// Create API client
api := ctx.NewAPI("https://api.example.com")

// Set default headers
api.SetHeader("X-API-Key", "your-key")

// Set auth token
api.SetAuth("your-jwt-token")

// Make requests
var users []User
api.Get(c, "/users", &users)

var newUser User
api.Post(c, "/users", payload, &newUser)

api.Put(c, "/users/123", updateData, &updatedUser)

api.Delete(c, "/users/123", nil)
```

## Example: Todo App with API

```go
var api = ctx.NewAPI("https://api.example.com")

func TodoPage(c *forge.Context) ui.UI {
    // Load todos on first render
    if c.Get("loaded") == nil {
        var todos []Todo
        if err := api.Get(c, "/todos", &todos); err == nil {
            c.Set("todos", todos)
        }
        c.Set("loaded", true)
    }
    
    todos := c.Get("todos").([]Todo)
    
    return ui.Div(
        ui.H1(ui.T("Todos")),
        ui.Button(ui.T("Add")).OnClick(c, func(c *forge.Context) {
            newTodo := Todo{Title: c.String("input")}
            var created Todo
            if err := api.Post(c, "/todos", newTodo, &created); err == nil {
                todos := c.Get("todos").([]Todo)
                c.Set("todos", append(todos, created))
                c.Set("input", "")
            }
        }),
        // ... render todos
    )
}
```

## Error Handling

```go
ui.Button(ui.T("Load")).OnClick(c, func(c *forge.Context) {
    c.Set("loading", true)
    c.Set("error", "")
    
    var data Response
    if err := c.FetchJSON(url, &data); err != nil {
        c.Set("error", "Failed to load: " + err.Error())
        c.Set("loading", false)
        return
    }
    
    c.Set("data", data)
    c.Set("loading", false)
})

// In UI
func Page(c *forge.Context) ui.UI {
    if c.Bool("loading") {
        return ui.Spinner()
    }
    if err := c.String("error"); err != "" {
        return ui.Alert(err, "error")
    }
    // render data...
}
```

## Custom Headers

```go
headers := map[string]string{
    "Authorization": "Bearer " + token,
    "X-Custom":      "value",
}

data, err := c.Request("GET", url, nil, headers)
```
