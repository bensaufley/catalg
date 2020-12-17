package models

import (
	"context"
	"database/sql"
	"errors"

	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/auth"
	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/validation"
)

type User struct {
	Model
	Username         string         `json:"username" gorm:"size:32;uniqueIndex;not null"`
	Email            string         `json:"email" gorm:"size:128;uniqueIndex;not null"`
	PasswordDigest   sql.NullString `json:"-" gorm:"size:96"`
	Salt             sql.NullString `json:"-" gorm:"size:32"`
	ActivatedAt      sql.NullTime
	EmailConfirmedAt sql.NullTime
}

// SetPassword takes a password, validates, and if valid, sets
// the PasswordDigest and Salt on the given User.
func (u *User) SetPassword(password string) error {
	if err := validation.Password(password); err != nil {
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

func CreateUser(ctx context.Context, db *gorm.DB, u User, password string) (*User, error) {
	err := validation.CollectErrors(
		validation.Username(u.Username),
		validation.Password(password),
	)
	if err != nil {
		return nil, err
	}

	user := User{
		Username: u.Username,
		Email:    u.Email,
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	if tx := db.WithContext(ctx).Create(&user); tx.Error != nil {
		log.WithError(tx.Error).Error("error creating user")
		return nil, errors.New("there was an unexpected error creating the user")
	}

	return &user, nil
}

type UserUpdateParams struct {
	Username *string
	Email    *string
	Password *string
}

func UpdateUser(ctx context.Context, db *gorm.DB, uuid string, password string, toUpdate UserUpdateParams) (*User, error) {
	user := &User{}
	tx := db.First(&user, "uuid = ?", uuid)
	if tx.Error != nil {
		log.WithError(tx.Error).Error("error looking up user for update")
		return nil, errors.New("could not update user")
	}
	lg := log.WithField("user", user.Username)
	if !user.PasswordDigest.Valid || !user.Salt.Valid {
		lg.Warn("attempting to update user without password")
		return nil, errors.New("could not update user")
	}
	if ok := auth.ComparePassword(password, user.PasswordDigest.String, user.Salt.String); !ok {
		lg.Warn("incorrect password to update user")
		return nil, errors.New("could not update user")
	}
	updateParams := map[string]interface{}{}
	for key, val := range map[string]*string{"username": toUpdate.Username, "email": toUpdate.Email} {
		if val != nil {
			updateParams[key] = *val
		}
	}
	if toUpdate.Password != nil {
		if err := user.SetPassword(*toUpdate.Password); err != nil {
			lg.WithError(err).Warn("could not set password in UpdateUser")
			return nil, errors.New("could not update user")
		}
		updateParams["password_digest"] = user.PasswordDigest.String
		updateParams["salt"] = user.Salt.String
	}
	if len(updateParams) > 0 {
		if tx := db.Model(user).Updates(updateParams); tx.Error != nil {
			lg.WithError(tx.Error).Error("error updating user")
			return nil, errors.New("could not update user")
		}
	}
	return nil, nil
}
