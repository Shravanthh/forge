package config

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Config holds application configuration.
type Config struct {
	data map[string]string
	mu   sync.RWMutex
}

// New creates an empty config.
func New() *Config {
	return &Config{data: make(map[string]string)}
}

// Load loads config from environment and optional JSON file.
func Load(jsonPath string) *Config {
	c := New()

	// Load JSON file if exists
	if jsonPath != "" {
		if data, err := os.ReadFile(jsonPath); err == nil {
			var m map[string]any
			if json.Unmarshal(data, &m) == nil {
				c.loadMap("", m)
			}
		}
	}

	// Environment variables override JSON
	for _, env := range os.Environ() {
		if i := strings.Index(env, "="); i > 0 {
			c.data[env[:i]] = env[i+1:]
		}
	}

	return c
}

func (c *Config) loadMap(prefix string, m map[string]any) {
	for k, v := range m {
		key := k
		if prefix != "" {
			key = prefix + "_" + k
		}
		key = strings.ToUpper(key)

		switch val := v.(type) {
		case map[string]any:
			c.loadMap(key, val)
		case string:
			c.data[key] = val
		case float64:
			c.data[key] = strconv.FormatFloat(val, 'f', -1, 64)
		case bool:
			c.data[key] = strconv.FormatBool(val)
		}
	}
}

// Get returns a string value.
func (c *Config) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

// GetOr returns value or default.
func (c *Config) GetOr(key, def string) string {
	if v := c.Get(key); v != "" {
		return v
	}
	return def
}

// Int returns an int value.
func (c *Config) Int(key string) int {
	v, _ := strconv.Atoi(c.Get(key))
	return v
}

// IntOr returns int or default.
func (c *Config) IntOr(key string, def int) int {
	if v := c.Get(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}

// Bool returns a bool value.
func (c *Config) Bool(key string) bool {
	return c.Get(key) == "true" || c.Get(key) == "1"
}

// Set sets a value at runtime.
func (c *Config) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
