package main

import (
	"fmt"
	"github.com/chicken-afk/go-fiber/database"
	"github.com/chicken-afk/go-fiber/database/migration"
	"github.com/chicken-afk/go-fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//Init Database
	database.DatabaseInit()
	migration.RunMigration()

	// Fiber instance
	app := fiber.New()

	//Init Route
	routes.RouteInit(app)

	// start server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
