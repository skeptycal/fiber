package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
)

func main() {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	if err != nil {
		log.Info(err)
	}

	fmt.Println(orgs)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/mike", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mike!")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
