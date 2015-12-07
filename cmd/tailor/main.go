package main

import (
	"golang.org/x/net/websocket"
	"log"
)

func main() {
	client, err := websocket.Dial("ws://localhost:9000/ws", "", "http://localhost/")
	if err != nil {
		log.Println(err)
		return
	}

	client.Write([]byte("Hello!"))
	for {
		msg := make([]byte, 512)
		if n, err := client.Read(msg); err != nil {
			log.Println(err)
			client.Close()
			return
		} else if len(msg) == 0 {
			continue
		} else {
			m := string(msg[0:n])
			log.Println(m)
		}
	}
}
