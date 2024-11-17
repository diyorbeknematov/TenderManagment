package handler

import (
	"fmt"
	"net/http"

	ws "tender/pkg/websocket"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(manager *ws.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("WebSocket Upgrade Error:", err)
			return
		}

		manager.AddClient(conn)
		defer func() {
			manager.RemoveClient(conn)
			conn.Close()
		}()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Read Error:", err)
				break
			}
		}
	}
}
