package routes

import (
	"github.com/Watson-Sei/anichat/api_v1/controllers"
	"github.com/Watson-Sei/anichat/api_v1/middleware/firebase"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App)  {
	admin := app.Group("/admin")

	admin.Use(firebase.Auth())

	admin.Get("/users", controllers.GetUsers)
}