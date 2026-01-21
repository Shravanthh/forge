# Configuration

Load config from environment variables and JSON files.

## Usage

```go
import "github.com/Shravanthh/forge/config"

// Load from env vars only
cfg := config.New()

// Load from JSON file + env vars (env overrides JSON)
cfg := config.Load("config.json")

port := cfg.GetOr("PORT", "3000")
debug := cfg.Bool("DEBUG")
```

## Config File

`config.json`:
```json
{
  "port": "3000",
  "debug": "true",
  "database": {
    "host": "localhost",
    "port": "5432"
  }
}
```

Nested keys are flattened with underscores:
```go
cfg.Get("DATABASE_HOST") // "localhost"
cfg.Get("DATABASE_PORT") // "5432"
```

## Environment Override

Environment variables override JSON values:

```bash
export PORT=8080
export DATABASE_HOST=prod-db.example.com
```

```go
cfg := config.Load("config.json")
cfg.Get("PORT")          // "8080" (from env)
cfg.Get("DATABASE_HOST") // "prod-db.example.com" (from env)
```

## API

### String Values

```go
cfg.Get("KEY")              // Returns "" if not set
cfg.GetOr("KEY", "default") // Returns "default" if not set
```

### Integer Values

```go
cfg.Int("PORT")           // Returns 0 if not set or invalid
cfg.IntOr("PORT", 3000)   // Returns 3000 if not set
```

### Boolean Values

```go
cfg.Bool("DEBUG") // true if value is "true" or "1"
```

### Runtime Updates

```go
cfg.Set("KEY", "value")
```

## Example

```go
package main

import (
    "github.com/Shravanthh/forge"
    "github.com/Shravanthh/forge/config"
    "github.com/Shravanthh/forge/log"
)

func main() {
    cfg := config.Load("config.json")
    
    logger := log.Default()
    if cfg.Bool("DEBUG") {
        logger.SetLevel(log.DEBUG)
    }
    
    app := forge.New()
    app.Route("/", HomePage)
    
    port := cfg.GetOr("PORT", "3000")
    logger.Info("Starting on port %s", port)
    app.Run(":" + port)
}
```

## Best Practices

1. Use env vars for secrets (never commit to JSON)
2. Use JSON for defaults and non-sensitive config
3. Use `GetOr`/`IntOr` to provide sensible defaults
4. Keep config keys uppercase for consistency
