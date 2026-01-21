# Logging

Structured JSON logging with levels.

## Usage

```go
import "github.com/Shravanthh/forge/log"

logger := log.Default() // INFO level
logger.Info("Server started on port %d", 3000)
```

Output:
```json
{"time":"2024-01-15T10:30:00Z","level":"INFO","msg":"Server started on port 3000"}
```

## Log Levels

```go
logger := log.New(log.DEBUG) // DEBUG, INFO, WARN, ERROR

logger.Debug("Processing request %s", id)
logger.Info("User logged in")
logger.Warn("Rate limit approaching")
logger.Error("Database connection failed: %v", err)
```

## Structured Fields

Add context to all log entries:

```go
// Add fields
logger := log.Default().
    With("service", "api").
    With("version", "1.0.0")

logger.Info("Request received")
// {"time":"...","level":"INFO","msg":"Request received","service":"api","version":"1.0.0"}

// Per-request context
reqLogger := logger.With("request_id", requestID)
reqLogger.Info("Processing")
```

## Change Level at Runtime

```go
logger.SetLevel(log.DEBUG) // Enable debug logs
logger.SetLevel(log.ERROR) // Only errors
```

## Example

```go
package main

import "github.com/Shravanthh/forge/log"

func main() {
    logger := log.New(log.INFO).With("app", "myapp")
    
    logger.Info("Starting server")
    
    // In request handler
    reqLog := logger.With("path", "/users")
    reqLog.Debug("Fetching users") // Won't print (below INFO)
    reqLog.Info("Found %d users", 42)
}
```
