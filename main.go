package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort = ":3000"
)

func main() {
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	if err != nil {
		log.Info(err)
	}

	fmt.Println(orgs)

	// err = GoFiber()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

// GoFiber runs the Fiber server listening on defaultPort
func GoFiber() error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/mike", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mike!")
	})

	err := app.Listen(defaultPort)
	if err != nil {
		return err
	}

	return nil
}
