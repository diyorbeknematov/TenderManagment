package handler

import (
	"fmt"
	"net/http"
	"tender/model"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin (you may want to restrict this in production)
		return true
	},
}

func (h *Handler) WebSocketNotifications(c *gin.Context) {
	// Retrieve the UserID from the context
	userID, exists := c.Get("UserID")
	if !exists {
		h.Log.Error("UserID not found in context")
		c.JSON(model.ErrUnauthorized.Code, model.ErrUnauthorized)
		return
	}

	// Upgrade the connection to a WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		h.Log.Error(fmt.Sprintf("Failed to upgrade connection: %v", err))
		return
	}
	defer conn.Close()

	// Fetch unread notifications for the user
	unreadNotifications, err := h.Service.GetAllNotifications(model.NotifFilter{UserID: userID.(string), IsRead: "false"})
	if err != nil {
		h.Log.Error(fmt.Sprintf("Error fetching unread notifications: %v", err))
		return
	}

	// Send unread notifications to the WebSocket connection
	for _, notif := range unreadNotifications.Notifications {
		if err := conn.WriteJSON(notif); err != nil {
			h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
			return
		}
	}

	// Optionally, you can handle incoming messages from the client
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			h.Log.Error(fmt.Sprintf("Error reading message: %v", err))
			break
		}
		h.Log.Info(fmt.Sprintf("Received message from client: %s", msg))
		// Handle the message as needed
	}
}
