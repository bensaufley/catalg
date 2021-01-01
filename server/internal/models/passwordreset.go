package models

import (
	"context"

	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/mailer"
	"github.com/bensaufley/catalg/server/internal/stubbables"
)

type PasswordReset struct {
	Model
	UserUUID string `gorm:"uniqueIndex;not null"`
	Token    string `gorm:"size:128;not null"`
}

func GeneratePasswordReset(ctx context.Context, db *gorm.DB, email string, client mailer.EmailClient) {
	user := &User{}
	if tx := db.WithContext(ctx).First(user, "email = ?", email); tx.Error != nil {
		log.WithError(tx.Error).WithField("email", email).Warn("error looking up user for password reset")
		return
	}
	lg := log.WithField("user", user.Username)
	token := stubbables.RandomChars(128)
	pr := PasswordReset{
		Token: token,
	}
	if err := db.Model(user).Association("PasswordReset").Append(&pr); err != nil {
		lg.WithError(err).Error("error creating PasswordReset for user")
	}

	if err := client.PasswordReset(ctx, mailer.PasswordResetEmailParams{
		Email:    user.Email,
		Username: user.Username,
		Token:    pr.Token,
	}); err != nil {
		lg.WithError(err).Error("error sending password reset email")
	}
}
