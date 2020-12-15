package main

import (
	"log"

    "github.com/gofiber/fiber/v2"
    gh "github.com/google/go-github/v33/github"
)

func main() {
    app := fiber.New()
    f  gh.ArchiveFormat = ""

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
