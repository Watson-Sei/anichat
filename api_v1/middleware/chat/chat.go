package chat

import (
	"github.com/Watson-Sei/anichat/api_v1/database"
	"github.com/Watson-Sei/anichat/api_v1/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"log"
)

func PublicChecker() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roomId := c.Path()[4:]
		log.Printf("roomId: %s", roomId)
		db := database.DBConn

		var room models.Room
		if err := db.Where("ID = ?", roomId).First(&room).Error; err != nil {
			log.Printf("ID 1 was not found in the Room table.")
			return err
		}

		if !room.Public {
			log.Println("The Public Boolean value for Room ID 1 was false.")
			return nil
		}

		c.Next()
		return nil
	}
}

func WebSocket() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	}
}
