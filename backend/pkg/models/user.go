package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primarykey" json:"id"`
	Name  string `json:"name"`
	Email string `gorm:"size:255;index:idx_email,unique" json:"email"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserPayload struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}
