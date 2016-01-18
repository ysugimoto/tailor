package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

// Create Websocket handler function with READER mode
func createReader(app *AppHandler) websocket.Handler {
	return func(conn *websocket.Conn) {
		fmt.Println("Reader Connection")
		c := NewConnection(READER, conn)
		c.OnClose = func() {
			delete(app.connections, c.Id)
		}
		c.OnMessage = func(message string) {
			// noop
		}

		app.connections[c.Id] = c
		c.Poll()
	}
}
