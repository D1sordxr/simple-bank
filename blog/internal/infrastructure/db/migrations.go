package db

import (
	"LearningArch/blog/internal/infrastructure/db/models"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.Base{}, &models.Author{}, &models.Post{}, &models.Tag{}); err != nil {
		panic(err.Error())
	}
}
