package server

import (
	"log"
	"net/http"
	"time"
)

// Middleware is a function that wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// Use adds middleware to the app.
func (a *App) Use(mw Middleware) {
	a.middleware = append(a.middleware, mw)
}

// Logger logs requests.
func Logger() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
		})
	}
}

// CORS adds CORS headers.
func CORS(origins ...string) Middleware {
	origin := "*"
	if len(origins) > 0 {
		origin = origins[0]
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Recover recovers from panics.
func Recover() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("panic: %v", err)
					http.Error(w, "Internal Server Error", 500)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// RateLimit limits requests per IP.
func RateLimit(requestsPerMinute int) Middleware {
	visitors := make(map[string]*rateLimiter)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			if rl, ok := visitors[ip]; ok {
				if !rl.allow() {
					http.Error(w, "Too Many Requests", 429)
					return
				}
			} else {
				visitors[ip] = newRateLimiter(requestsPerMinute)
			}
			next.ServeHTTP(w, r)
		})
	}
}

type rateLimiter struct {
	tokens    int
	max       int
	lastReset time.Time
}

func newRateLimiter(rpm int) *rateLimiter {
	return &rateLimiter{tokens: rpm, max: rpm, lastReset: time.Now()}
}

func (r *rateLimiter) allow() bool {
	if time.Since(r.lastReset) > time.Minute {
		r.tokens = r.max
		r.lastReset = time.Now()
	}
	if r.tokens > 0 {
		r.tokens--
		return true
	}
	return false
}
