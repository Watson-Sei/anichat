package routes

import (
	"github.com/Watson-Sei/anichat/api_v1/controllers"
	"github.com/Watson-Sei/anichat/api_v1/middleware/firebase"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {

	app.Get("/rooms", controllers.GetAllRooms)

	admin := app.Group("/admin")

	admin.Use(firebase.Auth())

	// User management Routes
	admin.Get("/users", controllers.GetUsers)
	admin.Put("/users", controllers.UpdateUser)

	// Room management Routes
	admin.Post("/rooms", controllers.RegisterRoom)
	admin.Put("/rooms/:id", controllers.UpdateRoom)
	admin.Delete("/rooms/:id", controllers.DeleteRoom)
}
