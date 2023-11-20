package controllers

import (
	"github.com/chicken-afk/go-fiber/database"
	"github.com/chicken-afk/go-fiber/pkg/models"
	"github.com/chicken-afk/go-fiber/pkg/request"
	"github.com/chicken-afk/go-fiber/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (r *AuthController) Login(c *fiber.Ctx) error {
	var user models.User

	LoginRequest := new(request.LoginRequest)

	c.BodyParser(LoginRequest)

	validate := validator.New()
	errs := validate.Struct(LoginRequest)

	if errs != nil {
		validationErrors := []ErrorResponse{}
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			validationErrors = append(validationErrors, elem)
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Validation error",
			"error":   validationErrors,
		})
	}

	//Check Email
	if err := database.DB.Debug().Where("email", LoginRequest.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "email not found",
		})
	}
	//Check Password

	output := response.LoginResponse{
		TokenType: "Bearer",
		Token:     "FKJLKJdkneoieoufieo",
		User:      user,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login success",
		"data":    output,
	})
}
