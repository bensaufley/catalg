package migrations

import (
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/models"
)

func Perform(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
