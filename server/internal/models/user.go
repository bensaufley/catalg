package models

import (
	"database/sql"
	"errors"

	"github.com/bensaufley/catalg/server/internal/auth"
	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/validators"
)

type User struct {
	Model
	Username         string         `gorm:"size:32;uniqueIndex"`
	Email            string         `gorm:"size:128;uniqueIndex"`
	PasswordDigest   sql.NullString `gorm:"size:72"`
	Salt             sql.NullString `gorm:"size:32"`
	ActivatedAt      sql.NullTime
	EmailConfirmedAt sql.NullTime
}

// SetPassword takes a password, validates, and if valid, sets
// the PasswordDigest and Salt on the given User.
func (u *User) SetPassword(password string) error {
	if err := validators.Password(password); err != nil {
		return errors.New("password is not valid")
	}

	digest, salt := auth.HashPassword(password)

	u.PasswordDigest = sql.NullString{String: digest, Valid: true}
	u.Salt = sql.NullString{String: salt, Valid: true}

	return nil
}

func (u *User) Authenticate(password string) error {
	if !u.PasswordDigest.Valid || !u.Salt.Valid {
		log.WithField("user", u.Username).Warn("attempting to authenticate user without configured password")
		return errors.New("password is not valid")
	}
	if ok := auth.ComparePassword(password, u.PasswordDigest.String, u.Salt.String); !ok {
		log.WithField("user", u.Username).Warn("password comparison did not match")
		return errors.New("password is not valid")
	}
	return nil
}
