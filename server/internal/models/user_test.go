package models_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/models"
	"github.com/bensaufley/catalg/server/internal/testutils"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		db       *gorm.DB
		u        models.User
		password string
	}
	testCases := []struct {
		it      string
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				got, err := models.CreateUser(testCase.args.ctx, testCase.args.db, testCase.args.u, testCase.args.password)

				if testutils.AssertError(test, testCase.wantErr, err) {
					assert.Equal(test, testCase.want, got)
				}
			})
		}
	})
}

func TestUpdateUser(t *testing.T) {
	testCases := []struct {
		it      string
		ctx     context.Context
		db      *gorm.DB
		user    models.UpdateUserParams
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				got, err := models.UpdateUser(testCase.ctx, testCase.db, testCase.user)

				if testutils.AssertError(test, testCase.wantErr, err) {
					assert.Equal(test, testCase.want, got)
				}
			})
		}
	})
}

func TestUser_Authenticate(t *testing.T) {
	type args struct {
		password string
	}
	testCases := []struct {
		it      string
		user    *models.User
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				err := testCase.user.Authenticate(testCase.args.password)

				testutils.AssertError(test, testCase.wantErr, err)
			})
		}
	})
}

func TestUser_SetPassword(t *testing.T) {
	type fields struct {
		Model            models.Model
		Username         string
		Email            string
		PasswordDigest   sql.NullString
		Salt             sql.NullString
		ActivatedAt      sql.NullTime
		EmailConfirmedAt sql.NullTime
	}
	type args struct {
		password string
	}
	testCases := []struct {
		it      string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				u := &models.User{
					Model:            testCase.fields.Model,
					Username:         testCase.fields.Username,
					Email:            testCase.fields.Email,
					PasswordDigest:   testCase.fields.PasswordDigest,
					Salt:             testCase.fields.Salt,
					ActivatedAt:      testCase.fields.ActivatedAt,
					EmailConfirmedAt: testCase.fields.EmailConfirmedAt,
				}
				err := u.SetPassword(testCase.args.password)

				testutils.AssertError(test, testCase.wantErr, err)
			})
		}
	})
}

func TestValidatePasswordChars(t *testing.T) {
	type args struct {
		fl validator.FieldLevel
	}
	testCases := []struct {
		it   string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}

	t.Run("parallel group", func(g *testing.T) {
		for _, tc := range testCases {
			testCase := tc
			g.Run(testCase.it, func(test *testing.T) {
				test.Parallel()

				got := models.ValidatePasswordChars(testCase.args.fl)

				assert.Equal(test, testCase.want, got)
			})
		}
	})
}
