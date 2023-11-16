package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (r *BookController) Index(c *fiber.Ctx) error {
	data := Book{
		ID:          1,
		Title:       "Seratus Satu Kisah Menuju Surga",
		Description: "Deskripsi buku yang sangat bagus",
	}

	res := Response{
		Status:  "success",
		Message: "Berhasil memuat data",
		Data:    data,
	}

	return c.Status(http.StatusOK).JSON(res)
}
