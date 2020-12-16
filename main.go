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

	ctx := context.Background()
	client := github.NewClient(nil)
	data, resp, err := client.Organizations.List(ctx, "willnorris", nil)
	if err != nil {
		log.Info(err)
	}

	fmt.Println(data)
	fmt.Println("====================================")
	fmt.Println(resp)

	lic, resp, err := client.Licenses.Get(ctx, "MIT")
	if err != nil {
		log.Info(err)
	}

	fmt.Println(lic)
	fmt.Println("====================================")
	fmt.Println(resp)

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
