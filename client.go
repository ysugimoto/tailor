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
	url := fmt.Sprintf("ws://%s", clientHost)
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
		var msg string
		if err := websocket.Message.Receive(c.conn, &msg); err == nil {
			fmt.Println(msg)
		} else {
			break
		}
	}
}
