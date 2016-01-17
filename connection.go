package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

const (
	READER = iota
	PROXY
)

type Connection struct {
	Type      int
	Id        string
	conn      *websocket.Conn
	OnClose   func()
	OnMessage func(string)
}

func NewConnection(connType int, conn *websocket.Conn) *Connection {
	hash := sha1.Sum([]byte(fmt.Sprint(time.Now().UnixNano())))

	return &Connection{
		Type: connType,
		Id:   hex.EncodeToString(hash[:len(hash)]),
		conn: conn,
	}
}

func (c *Connection) Send(message string) {
	if err := websocket.Message.Send(c.conn, message); err != nil {
		fmt.Println("[Error]", "Send message for ID", c.Id, err.Error())
	} else {
		fmt.Println("Sended", message)
	}
}

func (c *Connection) Poll() {
	var msg string
	for {
		if err := websocket.Message.Receive(c.conn, &msg); err != nil {
			break
		} else {
		}
	}
}
