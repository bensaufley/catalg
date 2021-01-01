package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/mailer"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB          *gorm.DB
	EmailClient mailer.EmailClient
}
