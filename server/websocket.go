package server

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/Shravanthh/forge/ctx"
	"github.com/Shravanthh/forge/diff"
	"github.com/Shravanthh/forge/render"
	"github.com/Shravanthh/forge/ui"

	"github.com/gorilla/websocket"
)

// PageFunc renders a page given context.
type PageFunc func(*ctx.Context) ui.UI

// Session holds connection state.
type Session struct {
	ID      string
	Conn    *websocket.Conn
	Context *ctx.Context
	LastUI  ui.UI
	Page    PageFunc
	mu      sync.Mutex
}

// Message from client.
type Message struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Response to client.
type Response struct {
	Type    string       `json:"type"`
	Patches []diff.Patch `json:"patches,omitempty"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// SessionManager manages active sessions.
type SessionManager struct {
	mu       sync.RWMutex
	sessions map[string]*Session
	store    ctx.SessionStore
}

// NewSessionManager creates a session manager.
func NewSessionManager(store ctx.SessionStore) *SessionManager {
	if store == nil {
		store = ctx.NewMemoryStore()
	}
	return &SessionManager{sessions: make(map[string]*Session), store: store}
}

var sessionCounter uint64

func generateSessionID() string {
	return "s" + uitoa(atomic.AddUint64(&sessionCounter, 1))
}

func uitoa(i uint64) string {
	if i == 0 {
		return "0"
	}
	var b [20]byte
	n := len(b)
	for i > 0 {
		n--
		b[n] = byte('0' + i%10)
		i /= 10
	}
	return string(b[n:])
}

// HandleWebSocket handles WebSocket connections.
func (sm *SessionManager) HandleWebSocket(page PageFunc, params map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("upgrade error: %v", err)
			return
		}

		registerDevClient(conn)
		defer unregisterDevClient(conn)

		sessionID := r.URL.Query().Get("session")
		if sessionID == "" {
			sessionID = generateSessionID()
		}

		c := ctx.New()
		c.Params = params
		if state, _ := sm.store.Load(sessionID); state != nil {
			c.RestoreState(state)
		}

		ui.ResetEventCounter()
		initialUI := page(c)

		session := &Session{
			ID:      sessionID,
			Conn:    conn,
			Context: c,
			LastUI:  initialUI,
			Page:    page,
		}

		sm.mu.Lock()
		sm.sessions[sessionID] = session
		sm.mu.Unlock()

		defer func() {
			sm.store.Save(sessionID, c.PersistentState())
			conn.Close()
			sm.mu.Lock()
			delete(sm.sessions, sessionID)
			sm.mu.Unlock()
		}()

		conn.WriteJSON(map[string]string{"type": "session", "id": sessionID})

		for {
			var msg Message
			if err := conn.ReadJSON(&msg); err != nil {
				break
			}
			if msg.Type == "event" {
				sm.handleEvent(session, msg)
			}
		}
	}
}

func (sm *SessionManager) handleEvent(s *Session, msg Message) {
	s.mu.Lock()
	defer s.mu.Unlock()

	func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("handler panic: %v", r)
			}
		}()
		if msg.Value != "" {
			s.Context.HandleWithValue(msg.ID, msg.Value)
		} else {
			s.Context.Handle(msg.ID)
		}
	}()

	ui.ResetEventCounter()
	newUI := s.Page(s.Context)
	patches := diff.Diff(s.LastUI, newUI)
	s.LastUI = newUI

	if len(patches) > 0 {
		s.Conn.WriteJSON(Response{Type: "patch", Patches: patches})
	}
}

// RenderInitialHTML renders the initial page HTML.
func RenderInitialHTML(page PageFunc) string {
	c := ctx.New()
	ui.ResetEventCounter()
	return render.HTML(page(c))
}
