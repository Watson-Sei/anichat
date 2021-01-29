package main

import (
	"github.com/gofiber/websocket/v2"
	"log"
)

type message struct {
	data []byte
	room string
}

type subscription struct {
	conn *websocket.Conn
	room string
}

type hub struct {
	rooms map[string]map[*websocket.Conn]bool
	broadcast chan message
	register chan subscription
	unregister chan subscription
}

var h = hub{
	rooms: make(map[string]map[*websocket.Conn]bool),
	broadcast: make(chan message),
	register: make(chan subscription),
	unregister: make(chan subscription),
}

func (h *hub) run() {
	for {
		select {
		case s := <- h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*websocket.Conn]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
			log.Println("connection registered")

		case s := <- h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					log.Println("connection unregistered")
				}
			}

		case m := <- h.broadcast:
			log.Println("message received:", m)
			connections := h.rooms[m.room]

			for connection := range connections {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(m.data)); err != nil {
					log.Println("write error:", err)

					delete(connections, connection)
					if len(connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}