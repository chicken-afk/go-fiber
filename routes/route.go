package routes

import (
	"github.com/chicken-afk/go-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	//Prefix
	api := route.Group("/api")
	v1 := api.Group("/v1")
	// Routes
	UserController := controllers.NewUserController()
	v1.Get("/user", UserController.Index)
	v1.Post("/user", UserController.Create)
	v1.Get("/user/:id", UserController.Detail)
	v1.Put("user/:id", UserController.Update)
	v1.Delete("user/:id", UserController.DeleteUser)

	BookController := controllers.NewBookController()
	v1.Get("/book", BookController.Index)
}
