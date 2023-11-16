package routes

import (
	"github.com/chicken-afk/go-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	// Routes
	UserController := controllers.NewUserController()
	route.Get("/user", UserController.Index)
	route.Post("/user", UserController.Create)

	BookController := controllers.NewBookController()
	route.Get("/book", BookController.Index)
}
