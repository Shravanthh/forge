package log

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// Level represents log severity.
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

var levelNames = []string{"DEBUG", "INFO", "WARN", "ERROR"}

// Logger provides structured logging.
type Logger struct {
	level  Level
	fields map[string]any
	mu     sync.Mutex
}

// New creates a logger with minimum level.
func New(level Level) *Logger {
	return &Logger{level: level, fields: make(map[string]any)}
}

// Default returns an INFO-level logger.
func Default() *Logger { return New(INFO) }

// With returns a logger with additional fields.
func (l *Logger) With(key string, value any) *Logger {
	newFields := make(map[string]any, len(l.fields)+1)
	for k, v := range l.fields {
		newFields[k] = v
	}
	newFields[key] = value
	return &Logger{level: l.level, fields: newFields}
}

// SetLevel changes minimum log level.
func (l *Logger) SetLevel(level Level) { l.level = level }

func (l *Logger) log(level Level, msg string, args ...any) {
	if level < l.level {
		return
	}

	entry := map[string]any{
		"time":  time.Now().Format(time.RFC3339),
		"level": levelNames[level],
		"msg":   fmt.Sprintf(msg, args...),
	}
	for k, v := range l.fields {
		entry[k] = v
	}

	l.mu.Lock()
	defer l.mu.Unlock()
	json.NewEncoder(os.Stdout).Encode(entry)
}

func (l *Logger) Debug(msg string, args ...any) { l.log(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...any)  { l.log(INFO, msg, args...) }
func (l *Logger) Warn(msg string, args ...any)  { l.log(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.log(ERROR, msg, args...) }
