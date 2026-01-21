# Deployment

Deploy Forge applications to production.

## Build for Production

```bash
forge build
```

This creates an optimized binary.

## Deployment Options

### 1. Direct Server

```bash
# Build
forge build

# Copy to server
scp ./my-app user@server:/app/

# Run on server
ssh user@server
cd /app
./my-app
```

### 2. Docker

```dockerfile
# Dockerfile
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -ldflags="-s -w" -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 3000
CMD ["./main"]
```

```bash
docker build -t my-forge-app .
docker run -p 3000:3000 my-forge-app
```

### 3. AWS EC2

```bash
# On EC2 instance
sudo yum install golang  # or apt-get
git clone https://github.com/user/my-app
cd my-app
go build -o app
./app
```

### 4. AWS ECS/Fargate

Use the Docker approach with ECS task definition.

### 5. Google Cloud Run

```bash
gcloud run deploy my-app \
  --source . \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

### 6. Fly.io

```toml
# fly.toml
app = "my-forge-app"

[build]
  builder = "paketobuildpacks/builder:base"

[[services]]
  internal_port = 3000
  protocol = "tcp"

  [[services.ports]]
    port = 80
    handlers = ["http"]

  [[services.ports]]
    port = 443
    handlers = ["tls", "http"]
```

```bash
fly launch
fly deploy
```

### 7. Railway

Connect GitHub repo, Railway auto-detects Go and deploys.

### 8. Render

```yaml
# render.yaml
services:
  - type: web
    name: my-forge-app
    env: go
    buildCommand: go build -o app
    startCommand: ./app
```

## Environment Variables

```go
port := os.Getenv("PORT")
if port == "" {
    port = "3000"
}
app.Run(":" + port)
```

## Reverse Proxy (Nginx)

```nginx
server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## SSL/TLS

### With Nginx + Let's Encrypt

```bash
sudo certbot --nginx -d example.com
```

### Built-in TLS

```go
app.RunTLS(":443", "cert.pem", "key.pem")
```

## Health Checks

Add a health endpoint:

```go
app.Route("/health", func(c *forge.Context) ui.UI {
    return ui.Raw{HTML: "OK"}
})
```

## Scaling Considerations

### WebSocket Sticky Sessions

Forge uses WebSocket, so you need sticky sessions for load balancing:

**AWS ALB:**
- Enable sticky sessions on target group

**Nginx:**
```nginx
upstream forge {
    ip_hash;  # sticky sessions
    server localhost:3001;
    server localhost:3002;
}
```

### Session Storage

For multi-instance deployments, use Redis for session storage:

```go
// Implement ctx.SessionStore interface
type RedisStore struct {
    client *redis.Client
}

func (r *RedisStore) Save(id string, state map[string]any) error {
    data, _ := json.Marshal(state)
    return r.client.Set(ctx, "session:"+id, data, 24*time.Hour).Err()
}

func (r *RedisStore) Load(id string) (map[string]any, error) {
    data, err := r.client.Get(ctx, "session:"+id).Bytes()
    if err != nil {
        return nil, err
    }
    var state map[string]any
    json.Unmarshal(data, &state)
    return state, nil
}
```

## Monitoring

### Logging

```go
import "log/slog"

slog.Info("server started", "port", 3000)
```

### Metrics

Add Prometheus metrics endpoint if needed.

## Production Checklist

- [ ] Build with `-ldflags="-s -w"` for smaller binary
- [ ] Set up reverse proxy (Nginx/Caddy)
- [ ] Configure SSL/TLS
- [ ] Set up health checks
- [ ] Configure sticky sessions for load balancing
- [ ] Use Redis for session storage (multi-instance)
- [ ] Set up logging and monitoring
- [ ] Configure firewall rules
- [ ] Set up automatic restarts (systemd/supervisor)
