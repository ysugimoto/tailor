package main

import (
	"github.com/ysugimoto/husky"
)

func main() {
	app := husky.NewApp()
	app.Config.Set("host", "0.0.0.0")
	app.Config.Set("port", 9000)

	app.WebSocket("/ws", func(ws *husky.WebSocketDispatcher) {
		ws.OnMessage = func(message string) {
			println(message)
		}
		err := ws.Broadcast([]byte(ws.Id + " has connected"))
		if err != nil {
			println(err.Error())
		}
	})

	app.Serve()
}
