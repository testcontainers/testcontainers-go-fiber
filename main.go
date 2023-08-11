package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type MyApp struct {
	// The name of the app
	Name string
	// The version of the app
	Version string
	// The fiber app
	FiberApp *fiber.App
	// The database connection string for the users database. The application will need it to connect to the database,
	// reading it from the USERS_CONNECTION environment variable in production, or from the container in development.
	UsersConnection string
}

var App *MyApp = &MyApp{
	Name:            "my-app",
	Version:         "0.0.1",
	UsersConnection: os.Getenv("USERS_CONNECTION"),
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// register the fiber app
	App.FiberApp = app

	log.Fatal(app.Listen(":8000"))
}
