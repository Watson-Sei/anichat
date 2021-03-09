package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Firebase Admin SDK Setup
		opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			firebaseUninitialized(c)
			return err
		}

		auth, err := app.Auth(context.Background())
		if err != nil {
			log.Printf("error initializing app: %v\n", err)
			firebaseUninitialized(c)
			return err
		}

		authorizationHeader := c.Get("Authorization")
		if authorizationHeader == "" {
			log.Println("not value access_token")
			unauthorized(c)
			return nil
		}
		if authorizationHeader[:7] != "Bearer " {
			log.Println("not Bearer String")
			unauthorized(c)
			return nil
		}
		tokenString := authorizationHeader[7:]

		// JWT Token auth
		token, err := auth.VerifyIDToken(context.Background(), tokenString)
		if err != nil {
			log.Printf("error verifying ID token: %v\n", err)
			unauthorized(c)
			return err
		}

		log.Printf("Verifield ID token %v\n", token)
		c.Set("uid", token.UID)
		c.Next()
		return nil
	}
}

func unauthorized(c *fiber.Ctx) {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": http.StatusText(http.StatusInternalServerError),
	})
}

func firebaseUninitialized(c *fiber.Ctx) {
	c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": http.StatusText(http.StatusInternalServerError),
	})
}
