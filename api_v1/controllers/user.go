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