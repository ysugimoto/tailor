package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

// WebSoket connection wrapper struct
// manage connection, and polling
type Client struct {
	// WebSocket connection
	conn *websocket.Conn
}

// Create New Client
func NewClient(clientHost string) (*Client, error) {
	url := fmt.Sprintf("ws://%s/reader", clientHost)
	origin := fmt.Sprintf("http://%s", clientHost)

	if conn, err := websocket.Dial(url, "", origin); err != nil {
		return nil, err
	} else {
		return &Client{
			conn: conn,
		}, nil
	}
}

// Connection polling
func (c *Client) Listen() {
	for {
		var p Payload
		if err := websocket.JSON.Receive(c.conn, &p); err == nil {
			fmt.Println(p.Message)
		} else {
			c.conn.Close()
			break
		}
	}
}
