package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

type Payload struct {
	Message string `json:"message"`
	Host    string `json:"host"`
}

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
	OnMessage func(Payload)
}

// Create New connection
func NewConnection(connType int, conn *websocket.Conn) *Connection {
	hash := sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano())))

	return &Connection{
		Type:      connType,
		Id:        hex.EncodeToString(hash[:len(hash)]),
		conn:      conn,
		OnClose:   func() {},
		OnMessage: func(p Payload) {},
	}
}

// Send message to this connection
func (c *Connection) Send(p Payload) {
	if err := websocket.JSON.Send(c.conn, p); err != nil {
		c.OnClose()
		c.conn.Close()
		//fmt.Println("[Error]", "Send message for ID", c.Id, err.Error())
	} else {
	}
}

// Connection polling
func (c *Connection) Poll() {
	for {
		var p Payload
		if err := websocket.JSON.Receive(c.conn, &p); err == nil {
			c.OnMessage(p)
		} else {
			c.OnClose()
			c.conn.Close()
		}
	}
}
