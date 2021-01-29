package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"log"
)

func main()  {

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	socketapp := app.Group("/ws")

	go h.run()

	socketapp.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	socketapp.Get("/:roomId", websocket.New(func(c *websocket.Conn) {
		roomId := c.Params("roomId")

		s := subscription{c, roomId}

		defer func() {
			h.unregister <- s
			s.conn.Close()
		}()

		h.register <- s

		for {
			messageType, msg, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error: ", err)
				}
				return
			}

			if messageType == websocket.TextMessage {
				m := message{msg, s.room}
				h.broadcast <- m
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	app.Listen(":8080")
}