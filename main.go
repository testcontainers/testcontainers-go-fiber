package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type DevDependency interface {
	Terminate(ctx context.Context) error
}

type MyApp struct {
	// The name of the app
	Name string
	// The version of the app
	Version string
	// The fiber app
	FiberApp *fiber.App
	// The dependencies for development mode
	DevDependencies []DevDependency
	// The database connection string for the users database. The application will need it to connect to the database,
	// reading it from the USERS_CONNECTION environment variable in production, or from the container in development.
	UsersConnection string
}

var App *MyApp = &MyApp{
	Name:            "my-app",
	Version:         "0.0.1",
	DevDependencies: []DevDependency{},
	UsersConnection: os.Getenv("USERS_CONNECTION"),
}

func main() {
	app := fiber.New()

	// helper function to stop the dependencies
	shutdownFn := func() error {
		ctx := context.Background()
		for _, dep := range App.DevDependencies {
			err := dep.Terminate(ctx)
			if err != nil {
				log.Println("Error terminating the backend dependency:", err)
				return err
			}
		}

		return nil
	}

	// register the shutdown function
	app.Hooks().OnShutdown(shutdownFn)

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		// also use the shutdown function when the SIGTERM or SIGINT signals are received
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\n", sig)
		err := shutdownFn()
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":8000"))
}
