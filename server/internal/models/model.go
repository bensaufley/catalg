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
func (m *Model) BeforeCreate(*gorm.DB) error {
	m.UUID = stubbables.UUIDv1()
	return nil
}

type updateMap map[string]interface{}

func (m *updateMap) assign(key string, val interface{}) *updateMap {
	(*m)[key] = val
	return m
}

func (m *updateMap) assignIfPresent(key string, val interface{}) *updateMap {
	if val == nil {
		return m
	}
	(*m)[key] = val
	return m
}
