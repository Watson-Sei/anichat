package routes

import (
	"fmt"
	"github.com/Watson-Sei/anichat/api_v1/controllers"
	"github.com/Watson-Sei/anichat/api_v1/middleware/firebase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
)

func SetupRouter(app *fiber.App) {

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: fmt.Sprintf("%s,%s", os.Getenv("AllowOrigin1"), os.Getenv("AllowOrigin2")),
	}))

	app.Get("/rooms", controllers.GetAllRooms)

	admin := app.Group("/admin")

	admin.Use(firebase.AuthAdmin())

	// User management Routes
	admin.Get("/users", controllers.GetUsers)
	admin.Put("/users", controllers.UpdateUser)

	// Room management Routes
	admin.Post("/rooms", controllers.RegisterRoom)
	admin.Put("/rooms/:id", controllers.UpdateRoom)
	admin.Delete("/rooms/:id", controllers.DeleteRoom)
}
