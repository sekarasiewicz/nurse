package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type DeletedAt sql.NullTime

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
}
