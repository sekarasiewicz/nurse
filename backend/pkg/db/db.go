package db

import (
	"github.com/sekarasiewicz/nurse/backend/pkg/config"
	"github.com/sekarasiewicz/nurse/backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}

func InitDatabase() *gorm.DB {
	db := GetDatabase()
	Migrate(db)

	return db
}
