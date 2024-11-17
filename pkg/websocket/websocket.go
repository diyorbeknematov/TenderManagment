package websocket

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	clients map[*websocket.Conn]bool
	lock    sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (m *Manager) AddClient(conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.clients[conn] = true
}

func (m *Manager) RemoveClient(conn *websocket.Conn) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.clients, conn)
}

func (m *Manager) BroadcastMessage(message string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	for conn := range m.clients {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Write Error:", err)
			conn.Close()
			delete(m.clients, conn)
		}
	}
}
