package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/models"
	"github.com/bensaufley/catalg/server/internal/testutils"
)

func TestModel_BeforeCreate(test *testing.T) {
	teardown := testutils.StubUUIDv1("my-test-uuid")
	defer teardown()

	m := &models.Model{}

	err := m.BeforeCreate(&gorm.DB{})

	assert.NoError(test, err)
	assert.Equal(test, "my-test-uuid", m.UUID)
}
