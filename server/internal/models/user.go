package models

import (
	"context"
	"database/sql"
	"errors"
	"regexp"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/auth"
	"github.com/bensaufley/catalg/server/internal/log"
)

// User is the core user model
type User struct {
	Model
	Username         string         `json:"username" gorm:"size:32;uniqueIndex;not null" validate:"required,min=8,max=32,alphanum"`
	Email            string         `json:"email" gorm:"size:128;uniqueIndex;not null" validate:"required,email,max=128"`
	PasswordDigest   sql.NullString `json:"-" gorm:"size:96"`
	Salt             sql.NullString `json:"-" gorm:"size:32"`
	ActivatedAt      sql.NullTime
	EmailConfirmedAt sql.NullTime

	PasswordReset PasswordReset `json:"-"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("passwordchars", ValidatePasswordChars); err != nil {
		log.WithError(err).Fatal("could not register custom validator passwordchars")
	}
}

// SetPassword takes a password, validates, and if valid, sets
// the PasswordDigest and Salt on the given User.
func (u *User) SetPassword(password string) error {
	if err := validate.Var(password, "required,min=12,max=128,passwordchars"); err != nil {
		return err
	}

	digest, salt := auth.HashPassword(password)

	u.PasswordDigest = sql.NullString{String: digest, Valid: true}
	u.Salt = sql.NullString{String: salt, Valid: true}

	return nil
}

// Authenticate authenticates a user with a given password
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

// CreateUser validates and creates a user
func CreateUser(ctx context.Context, db *gorm.DB, u User, password string) (*User, error) {
	user := User{
		Username: u.Username,
		Email:    u.Email,
	}
	if err := validate.Struct(user); err != nil {
		return nil, err
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

var userUpdateError = errors.New("could not update user")

func handleUpdateUserError(err error, msg string, lg *log.Entry) (*User, error) {
	l := lg
	if l == nil {
		l = log.WithError(err)
	} else {
		l = lg.WithError(err)
	}
	l.Error(msg)
	return nil, userUpdateError
}

func UpdateUser(ctx context.Context, db *gorm.DB, params UpdateUserParams) (*User, error) {
	user := &User{}
	if tx := db.WithContext(ctx).First(&user, "uuid = ? AND salt IS NOT NULL AND password_digest IS NOT NULL", params.UUID); tx.Error != nil {
		return handleUpdateUserError(tx.Error, "error looking up user for update", nil)
	}
	lg := log.WithField("user", user.Username)
	if ok := auth.ComparePassword(params.Password, user.PasswordDigest.String, user.Salt.String); !ok {
		lg.Warn("incorrect password to update user")
		return nil, userUpdateError
	}
	updateParams := &updateMap{}
	updateParams.assignIfPresent("username", *params.Username).assignIfPresent("email", *params.Email)
	if params.NewPassword != nil {
		if err := user.SetPassword(*params.NewPassword); err != nil {
			return handleUpdateUserError(err, "could not set password in UpdateUser", lg)
		}
		updateParams.assign("password_digest", user.PasswordDigest.String).assign("salt", user.Salt.String)
	}
	if len(*updateParams) > 0 {
		if tx := db.WithContext(ctx).Model(user).Updates(updateParams); tx.Error != nil {
			return handleUpdateUserError(tx.Error, "error updating user", lg)
		}
	}
	return nil, nil
}

var letterRegExp = regexp.MustCompile(`[A-Za-z]`)
var numSymRegExp = regexp.MustCompile(`[0-9-=[\]\\;',./~!@#$%^&*()_+{}|:"<>?]`)

func ValidatePasswordChars(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	return letterRegExp.MatchString(str) && numSymRegExp.MatchString(str)
}
