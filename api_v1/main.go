package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func main()  {

	go h.run()

	app := fiber.New()

	app.Static("/", "./home.html")

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			h.unregister <- c
			c.Close()
		}()

		h.register <- c

		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}
				return
			}

			if messageType == websocket.TextMessage {
				h.broadcast <- string(msg)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	log.Fatal(app.Listen(":8080"))
}