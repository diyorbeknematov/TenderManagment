package handler

import (
	"fmt"
	"net/http"
	"sync"
	"tender/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (consider restricting this in production)
	},
}

// WebSocketHub manages active WebSocket connections
type WebSocketHub struct {
	connections map[*websocket.Conn]bool
	mu          sync.Mutex
}

var hub = WebSocketHub{
	connections: make(map[*websocket.Conn]bool),
}

// Add connection to the hub
func (h *WebSocketHub) register(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.connections[conn] = true
}

// Remove connection from the hub
func (h *WebSocketHub) unregister(conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.connections, conn)
}

// Broadcast a notification to all connected clients
func (h *WebSocketHub) broadcast(notification model.Notification) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for conn := range h.connections {
		if err := conn.WriteJSON(notification); err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
			conn.Close()
			delete(h.connections, conn)
		} else {
			fmt.Printf("Notification sent to client: %v\n", notification)
		}
	}
}

func (h *Handler) WebSocketNotifications(c *gin.Context) {
	userID, exists := c.Get("UserID")
	if !exists {
		h.Log.Error("UserID not found in context")
		c.JSON(model.ErrUnauthorized.Code, model.ErrUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to upgrade connection: %v", err))
		return
	}
	defer conn.Close()

	// Log successful connection
	h.Log.Info(fmt.Sprintf("WebSocket connection established for user: %s", userID))

	// Register the connection
	hub.register(conn)
	defer hub.unregister(conn)

	// Fetch unread notifications for the user
	unreadNotifications, err := h.Service.GetAllNotifications(model.NotifFilter{UserID: userID.(string), IsRead: "false"})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error fetching unread notifications: %v", err))
		return
	}

	// Send unread notifications to the WebSocket connection
	for _, notif := range unreadNotifications {
		if err := conn.WriteJSON(notif); err != nil {
			h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
			return
		}
	}

	// Handle incoming messages from the client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.Log.Error(fmt.Sprintf("Error reading message: %v", err))
			break
		}
		h.Log.Info(fmt.Sprintf("Received message from client: %s", msg))
	}
}

// Function to create a new notification and broadcast it
func (h *Handler) CreateNotification(userID string, message string, ty string, id string) error {
	// Create the notification in the database
	notification := model.Notification{
		UserID:     userID,
		Message:    message,
		IsRead:     false,
		Type:       ty,
		RelationID: id,
	}

	_, err := h.Service.CreateNotification(notification) // Assuming this function exists
	if err != nil {
		return err
	}

	// Broadcast the new notification to all connected clients
	hub.broadcast(notification)

	return nil
}
