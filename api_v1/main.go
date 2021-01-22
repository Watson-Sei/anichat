package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

type client struct{}

var clients = make(map[*websocket.Conn]client)
var register = make(chan *websocket.Conn)
var unregister = make(chan *websocket.Conn)
var broadcast = make(chan string)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = client{}
			log.Println("connection registered")

		case message := <-broadcast:
			log.Println("message received:", message)

			// Send the message to all clients
			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
					log.Println("write error:", err)

					unregister <- connection
					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
				}
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func main()  {
	app := fiber.New()

	app.Static("/", "./home.html")

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	go runHub()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()

		register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				broadcast <- string(message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	log.Fatal(app.Listen(":8080"))
}