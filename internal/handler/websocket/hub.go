package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.RWMutex
	clients map[string]*websocket.Conn // userID → conn
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[string]*websocket.Conn),
	}
}

func (h *Hub) Register(userID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[userID] = conn
}

func (h *Hub) Unregister(userID string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	delete(h.clients, userID)
}

func (h *Hub) Get(userID string) (*websocket.Conn, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	conn, ok := h.clients[userID]
	return conn, ok
}
