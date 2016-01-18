package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

const (
	// READER type connection
	READER = iota
	// PROXY type connection
	PROXY
)

// WebSocket connection wrapper struct
// Manage connection-id, and close/message event handler.
type Connection struct {
	// Connection type
	Type int

	// Connection ID
	// this fields supplied after instantiate
	Id string

	// WebSocket connection
	conn *websocket.Conn

	// Event handler of close connection
	OnClose func()

	// Event handler of message incoming
	OnMessage func(string)
}

// Create New connection
func NewConnection(connType int, conn *websocket.Conn) *Connection {
	hash := sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano())))

	return &Connection{
		Type:      connType,
		Id:        hex.EncodeToString(hash[:len(hash)]),
		conn:      conn,
		OnClose:   func() {},
		OnMessage: func(message string) {},
	}
}

// Send message to this connection
func (c *Connection) Send(message string) {
	if err := websocket.Message.Send(c.conn, message); err != nil {
		fmt.Println("[Error]", "Send message for ID", c.Id, err.Error())
	} else {
		fmt.Println("Sended", message)
	}
}

// Connection polling
func (c *Connection) Poll() {
	for {
		var msg string
		if err := websocket.Message.Receive(c.conn, &msg); err == nil {
			c.OnMessage(msg)
		} else {
			break
		}
	}
}
