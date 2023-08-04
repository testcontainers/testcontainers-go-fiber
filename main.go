package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

type DevDependency interface {
	Terminate(ctx context.Context) error
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8000"))
}
