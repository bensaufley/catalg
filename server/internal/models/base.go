package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/stubbables"
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
	m.UUID = stubbables.UUIDv1()
	return nil
}
