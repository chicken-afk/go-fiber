package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Uuid      uuid.UUID      `json:"uuid"`
	Name      string         `json:"name"`
	Password  string         `json:"-" gorm:"column:password"`
	Email     string         `json:"email"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"column:deleted_at"`
}
