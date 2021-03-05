package controllers

import (
	"fmt"
	"github.com/Watson-Sei/anichat/api_v1/database"
	"github.com/Watson-Sei/anichat/api_v1/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

type ResponseHTTP struct {
	Success bool 		`json:"success"`
	Data 	interface{}	`json:"data"`
	Message string		`json:"message"`
}

func GetAllRooms(c *fiber.Ctx) error {
	db := database.DBConn

	var rooms []models.Room
	if res := db.Find(&rooms); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data: nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all rooms",
		Data: rooms,
	})
}

func RegisterRoom(c *fiber.Ctx) error {
	db := database.DBConn

	room := new(models.Room)
	if err := c.BodyParser(&room); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
	}

	db.Create(room)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a room.",
		Data: *room,
	})
}

func UpdateRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	afterRoom := new(models.Room)

	if err := c.BodyParser(afterRoom); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	beforeRoom := new(models.Room)
	if err := db.First(&beforeRoom, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Room with ID %v not found.", id),
				Data: nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data: nil,
			})
		}
	}

	db.Where("ID = ?", id).First(&beforeRoom).Updates(&afterRoom)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success update room",
		Data: *afterRoom,
	})
}

func DeleteRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	room := new(models.Room)
	if err := db.First(&room, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Room with ID %v not found.", id),
				Data: nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data: nil,
			})
		}
	}

	db.Delete(&room)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success delete room.",
		Data: nil,
	})
}