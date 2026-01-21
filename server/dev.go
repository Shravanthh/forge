package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	devClients   = make(map[*websocket.Conn]bool)
	devClientsMu sync.RWMutex
)

func registerDevClient(conn *websocket.Conn) {
	devClientsMu.Lock()
	devClients[conn] = true
	devClientsMu.Unlock()
}

func unregisterDevClient(conn *websocket.Conn) {
	devClientsMu.Lock()
	delete(devClients, conn)
	devClientsMu.Unlock()
}

// DevServer adds hot reload capability.
type DevServer struct {
	*App
	watchDir   string
	reloadChan chan struct{}
}

// NewDev creates a dev server with hot reload.
func NewDev(watchDir string) *DevServer {
	return &DevServer{
		App:        New(),
		watchDir:   watchDir,
		reloadChan: make(chan struct{}, 1),
	}
}

// Run starts dev server with file watching.
func (d *DevServer) Run(addr string) error {
	go d.watchFiles()
	go d.broadcastReloads()
	log.Printf("Forge DEV running at http://localhost%s (hot reload enabled)\n", addr)
	return http.ListenAndServe(addr, d)
}

func (d *DevServer) watchFiles() {
	lastMod := time.Now()
	for {
		time.Sleep(500 * time.Millisecond)
		changed := false
		filepath.Walk(d.watchDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			ext := filepath.Ext(path)
			if (ext == ".go" || ext == ".css") && info.ModTime().After(lastMod) {
				changed = true
				log.Printf("Changed: %s\n", path)
			}
			return nil
		})
		if changed {
			lastMod = time.Now()
			select {
			case d.reloadChan <- struct{}{}:
			default:
			}
		}
	}
}

func (d *DevServer) broadcastReloads() {
	for range d.reloadChan {
		devClientsMu.RLock()
		for conn := range devClients {
			conn.WriteJSON(map[string]string{"type": "reload"})
		}
		devClientsMu.RUnlock()
	}
}
