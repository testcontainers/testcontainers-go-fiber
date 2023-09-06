package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/testcontainers/testcontainers-go-fiber/db"
	"log"
	"os"
)

func main() {
	connStr := fmt.Sprintf(os.Getenv("DB_CONNECTION"))

	conn := db.GetDb(connStr)
	repo := db.NewUserRepo(conn)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		users, err := repo.GetUsers(c.UserContext())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":8080"))
}
