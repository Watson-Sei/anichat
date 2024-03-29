package main

import (
	"github.com/Watson-Sei/anichat/api_v1/database"
	"github.com/Watson-Sei/anichat/api_v1/middleware/chat"
	"github.com/Watson-Sei/anichat/api_v1/models"
	"github.com/Watson-Sei/anichat/api_v1/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
	"log"
)

var SecretKey string = "+_z#&=@+)^xpok3$#_@vg3xd$3avp8gj&_dx#9u-f(v+5lgs7@"

func main() {

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database:", err.Error())
	}

	database.DBConn.AutoMigrate(&models.Room{})

	routes.SetupRouter(app)

	socketapp := app.Group("/ws")

	go h.run()

	socketapp.Use(chat.PublicChecker(), chat.WebSocket())

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
				m := message{s.conn, msg, s.room}
				h.broadcast <- m
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	app.Listen(":8080")
}
