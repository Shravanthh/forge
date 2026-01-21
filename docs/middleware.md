# Middleware

Add middleware for logging, CORS, rate limiting, and more.

## Usage

```go
app := forge.New()

app.Use(server.Logger())
app.Use(server.CORS("*"))
app.Use(server.Recover())
app.Use(server.RateLimit(100))

app.Route("/", HomePage)
app.Run(":3000")
```

## Built-in Middleware

### Logger

Logs all requests with method, path, and duration.

```go
app.Use(server.Logger())
// Output: GET /users 1.234ms
```

### CORS

Adds CORS headers for cross-origin requests.

```go
// Allow all origins
app.Use(server.CORS("*"))

// Allow specific origin
app.Use(server.CORS("https://example.com"))
```

### Recover

Recovers from panics and returns 500 error.

```go
app.Use(server.Recover())
```

### Rate Limit

Limits requests per IP address.

```go
// 100 requests per minute per IP
app.Use(server.RateLimit(100))
```

## Custom Middleware

```go
func AuthMiddleware() server.Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            token := r.Header.Get("Authorization")
            if token == "" {
                http.Error(w, "Unauthorized", 401)
                return
            }
            // Validate token...
            next.ServeHTTP(w, r)
        })
    }
}

app.Use(AuthMiddleware())
```

## Middleware Order

Middleware executes in the order added:

```go
app.Use(server.Logger())    // 1st: logs request
app.Use(server.Recover())   // 2nd: catches panics
app.Use(server.CORS("*"))   // 3rd: adds headers
app.Use(AuthMiddleware())   // 4th: checks auth
```
