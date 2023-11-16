package migration

import (
	"fmt"
	"github.com/chicken-afk/go-fiber/database"
	"github.com/chicken-afk/go-fiber/pkg/models"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated Successfully")
}
