package controllers

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"os"
)

type UserInfo struct {
	UID string `json:"rawId"`
	DisplayName string `json:"displayName"`
	Email string `json:"email"`
}

func GetUsers(c *fiber.Ctx) error {
	ctx := context.Background()

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(fmt.Errorf("error initializing app: %v\n", err))
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	UserData := map[string][]*auth.UserInfo{}

	pager := iterator.NewPager(client.Users(ctx, ""), 100, "")
	for {
		var users []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			log.Fatalf("paging error %v\n", err)
		}
		for _, u := range users {
			UserData["users"] = append(UserData["users"], u.UserInfo)
			log.Printf("read user user: %v\n", u.UserInfo)
		}
		if nextPageToken == "" {
			break
		}
	}

	bytes, err := json.Marshal(UserData)
	if err != nil {
		log.Printf("error json encoding: %v\n", err)
	}

	return c.Status(fiber.StatusOK).JSON(string(bytes))
}

func UpdateUser(c *fiber.Ctx) error {
	p := new(UserInfo)

	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ctx := context.Background()

	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(fmt.Errorf("error initializing app: %v\n", err))
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToUpdate{}).
		Email(p.Email).
		DisplayName(p.DisplayName)
	u, err := client.UpdateUser(ctx, p.UID, params)

	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully fetched user data: %#v\n", u.UserInfo)

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Profile successfully update",
	}); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error updating profile",
		})
	}

	return nil
}