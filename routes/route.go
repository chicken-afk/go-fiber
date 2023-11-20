package routes

import (
	"github.com/chicken-afk/go-fiber/pkg/controllers"
	"github.com/chicken-afk/go-fiber/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(route *fiber.App) {
	//Prefix
	api := route.Group("/api")
	v1 := api.Group("/v1")

	//Auth
	AuthController := controllers.NewAuthController()
	v1.Get("/login", AuthController.Login)

	// Routes
	UserController := controllers.NewUserController()
	v1.Get("/user", middleware.AuthMiddleware, UserController.Index)
	v1.Post("/user", UserController.Create)
	v1.Get("/user/:id", UserController.Detail)
	v1.Put("user/:id", UserController.Update)
	v1.Delete("user/:id", UserController.DeleteUser)

	BookController := controllers.NewBookController()
	v1.Get("/book", BookController.Index)

}
