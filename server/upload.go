package server

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// UploadHandler handles file uploads.
type UploadHandler func(filename string, data []byte) error

// HandleUpload registers a file upload endpoint.
func (a *App) HandleUpload(path string, handler UploadHandler) {
	a.uploads[path] = handler
}

// ServeHTTP updated to handle uploads - see http.go

// SaveToDir returns an UploadHandler that saves files to a directory.
func SaveToDir(dir string) UploadHandler {
	return func(filename string, data []byte) error {
		os.MkdirAll(dir, 0755)
		return os.WriteFile(filepath.Join(dir, filename), data, 0644)
	}
}

func handleUpload(w http.ResponseWriter, r *http.Request, handler UploadHandler) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	r.ParseMultipartForm(32 << 20) // 32MB max
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := handler(header.Filename, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(`{"ok":true}`))
}
