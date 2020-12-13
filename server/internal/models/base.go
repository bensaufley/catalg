package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Model overrides the base Gorm model type to use UUID for id
type Model struct {
	UUID      string `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// BeforeCreate sets the UUID
func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.UUID = uuid.NewV1().String()
	return nil
}
