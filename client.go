package main

import (
	"fmt"
	"golang.org/x/net/websocket"
)

type Client struct {
	conn *websocket.Conn
}

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
