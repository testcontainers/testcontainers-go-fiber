package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := getDb()
	repo := repo{db: db}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/users", func(c *fiber.Ctx) error {
		users, err := repo.getUsers(c.UserContext())
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(users)
	})

	log.Fatal(app.Listen(":8080"))
}

type User struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type repo struct {
	db *pgx.Conn
}

func (r repo) getUsers(ctx context.Context) ([]User, error) {
	var users []User
	rows, err := r.db.Query(ctx, "select id, fullname,email from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user = User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getDb() *pgx.Conn {
	connStr := fmt.Sprintf(os.Getenv("DB_CONNECTION"))
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
