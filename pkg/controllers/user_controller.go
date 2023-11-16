package controllers

import (
	"github.com/chicken-afk/go-fiber/database"
	"github.com/chicken-afk/go-fiber/pkg/models"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Paginate struct {
	CurrentPage int         `json:"current_page"`
	PerPage     int         `json:"per_page"`
	TotalData   int64       `json:"total_data"`
	Data        interface{} `json:"data"`
}

func (r *UserController) Index(c *fiber.Ctx) error {
	var users []models.User
	var totalData int64

	// Menentukan ukuran halaman default
	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: "Gagal memuat data",
		})
	}

	// Mendapatkan nilai query parameter "page" (halaman) dan "pageSize" (ukuran halaman)
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("pageSize", strconv.Itoa(pageSize)))
	if err != nil || limit < 1 {
		limit = pageSize
	}

	// Menghitung offset untuk pagination
	offset := (page - 1) * limit

	search := c.Query("search", "")

	if search != "" {
		err = database.DB.Offset(offset).Limit(limit).Where("name LIKE ?", "%"+search+"%").Find(&users).Error
	} else {
		if err := database.DB.Table("users").Count(&totalData).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(Response{
				Status:  "error",
				Message: "Gagal memuat data",
			})
		}
		err = database.DB.Offset(offset).Limit(limit).Find(&users).Error
	}

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(Response{
			Status:  "error",
			Message: "Gagal memuat data",
		})
	}
	data := Paginate{
		CurrentPage: page,
		PerPage:     pageSize,
		TotalData:   totalData,
		Data:        users,
	}
	res := Response{
		Status:  "success",
		Message: "Berhasil memuat data",
		Data:    data,
	}

	return c.JSON(res)
}

func (r *UserController) Create(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Berhasil membuat user",
	})
}
