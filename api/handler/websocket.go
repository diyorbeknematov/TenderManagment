package handler

import (
	"fmt"
	"net/http"
	"tender/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (adjust as needed)
	},
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

	// Poll for unread notifications every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				unreadNotifications, err := h.Service.GetAllNotifications(model.NotifFilter{UserID: userID.(string), IsRead: "false"})
				if err != nil {
					h.Log.Error(fmt.Sprintf("Error fetching unread notifications: %v", err))
					continue
				}

				// Send unread notifications to the WebSocket connection
				for _, notif := range unreadNotifications {
					if err := conn.WriteJSON(notif); err != nil {
						h.Log.Error(fmt.Sprintf("Error sending notification: %v", err))
						return
					}
				}
			}
		}
	}()

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
