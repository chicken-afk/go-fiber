package controllers

import (
	"github.com/chicken-afk/go-fiber/database"
	"github.com/chicken-afk/go-fiber/pkg/models"
	"github.com/chicken-afk/go-fiber/pkg/request"
	"github.com/chicken-afk/go-fiber/pkg/response"
	"github.com/chicken-afk/go-fiber/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
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
	isValid := utils.CheckPasswordHash(LoginRequest.Password, user.Password)

	if !isValid {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "failed",
			"message": "wrong password",
		})
	}

	//Generate JWT
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Internal server error",
			"error":   err,
		})
	}

	output := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
		User:      user,
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login success",
		"data":    output,
	})
}
