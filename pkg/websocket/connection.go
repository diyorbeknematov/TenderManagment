package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Conn *websocket.Conn
}

func (c *Client) Send(message string) error {
	return c.Conn.WriteMessage(websocket.TextMessage, []byte(message))
}
