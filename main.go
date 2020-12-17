package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/v33/github"
	log "github.com/sirupsen/logrus"
)

const (
	defaultPort = 3000
)

func main() {
	err := GoFiber()
	if err != nil {
		log.Fatal(err)
	}
}

// GoFiber runs the Fiber server listening on the given port (or default)
func GoFiber() error {

	const sep = string(os.PathSeparator)

	home, _ := os.UserHomeDir()
	portPtr := flag.Int("port", defaultPort, "Port to listen on.")
	port := fmt.Sprintf(":%d", portPtr)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/mike", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Mike!")
	})

	userPicPath := filepath.Join(home, "Pictures")
	app.Static("/pictures", userPicPath)

	app.Get("/license", func(c *fiber.Ctx) error {
		return c.SendString(exampleLicense())
	})

	err := app.Listen(port)
	if err != nil {
		return err
	}

	return nil
}

func exampleLicense() string {
	ctx := context.Background()
	client := github.NewClient(nil)
	_, _, err := client.Organizations.List(ctx, "willnorris", nil)
	if err != nil {
		log.Info(err)
	}

	lic, _, _ := client.Licenses.Get(ctx, "MIT")
	if err != nil {
		log.Info(err)
	}

	return *lic.Body
}
