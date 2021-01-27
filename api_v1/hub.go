package main

import (
	"github.com/gofiber/websocket/v2"
	"log"
)

type client struct{}

type hub struct {
	clients map[*websocket.Conn]client
	register chan *websocket.Conn
	unregister chan *websocket.Conn
	broadcast chan string
}

var h = hub{
	clients: make(map[*websocket.Conn]client),
	register: make(chan *websocket.Conn),
	unregister: make(chan *websocket.Conn),
	broadcast: make(chan string),
}

func (h *hub) run() {
	for {
		select {
		case connection := <-h.register:
			h.clients[connection] = client{}
			log.Println("connection registered")

		case connection := <-h.unregister:
			delete(h.clients, connection)
			log.Println("connection unregistered")

		case message := <-h.broadcast:
			log.Println("message received:", message)
		}
	}
}