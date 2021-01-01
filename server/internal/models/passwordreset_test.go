package models_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/khaiql/dbcleaner.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/mailer"
	"github.com/bensaufley/catalg/server/internal/models"
	"github.com/bensaufley/catalg/server/internal/testutils"
)

var cleaner = dbcleaner.New()

type testCase struct {
	it        string
	setup     func(context.Context, *gorm.DB, *testing.T) func()
	email     string
	wantEmail *mailer.PasswordResetEmailParams
}

type PasswordResetSuite struct {
	suite.Suite
	cases []testCase

	db *gorm.DB
}

func (suite *PasswordResetSuite) SetupSuite() {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
	pg := testutils.NewPostgresEngine(dsn)
	cleaner.SetEngine(pg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: &log.GormLogger{}})
	if err != nil {
		panic(err)
	}
	suite.db = db

	suite.cases = []testCase{
		{
			it: "sends an email if everything is in order",

			setup: func(ctx context.Context, db *gorm.DB, test *testing.T) func() {
				user := &models.User{
					Email:    "successful@test.com",
					Username: "successfulPasswordResetUser",
				}
				if tx := db.WithContext(ctx).Create(user); tx.Error != nil {
					assert.FailNow(test, tx.Error.Error())
					return func() {}
				}
				return func() {
					if tx := db.WithContext(ctx).Unscoped().Select("PasswordReset").Delete(user); tx.Error != nil {
						assert.FailNow(test, tx.Error.Error())
					}
				}
			},
			email: "successful@test.com",
			wantEmail: &mailer.PasswordResetEmailParams{
				Email:    "successful@test.com",
				Username: "successfulPasswordResetUser",
				Title:    "",
				Token:    "XXXXXXXXXXXXXXXXXX",
			},
		},
	}
}

func (suite *PasswordResetSuite) TearDownSuite() {
	if err := cleaner.Close(); err != nil {
		panic(err)
	}
}

func (suite *PasswordResetSuite) SetupTest() {
	cleaner.Acquire("users", "password_resets")
}

func (suite *PasswordResetSuite) TearDownTest() {
	cleaner.Clean("users", "password_resets")
}

type testClient struct {
	calls []mailer.PasswordResetEmailParams
}

func (c *testClient) PasswordReset(_ context.Context, params mailer.PasswordResetEmailParams) error {
	p := params
	if p.Token != "" {
		p.Token = "XXXXXXXXXXXXXXXXXX"
	}
	c.calls = append(c.calls, p)
	return nil
}

func (c *testClient) assertCalledWithParams(test *testing.T, params mailer.PasswordResetEmailParams) bool {
	return assert.Contains(test, c.calls, params)
}

func (c *testClient) assertNotCalled(test *testing.T) bool {
	return assert.Empty(test, c.calls)
}

func (suite *PasswordResetSuite) TestPasswordReset() {
	for _, c := range suite.cases {
		testCase := c
		suite.Run(testCase.it, func() {
			suite.T().Parallel()

			ctx := context.Background()

			teardown := testCase.setup(ctx, suite.db, suite.T())
			defer teardown()
			client := &testClient{}

			models.GeneratePasswordReset(ctx, suite.db, testCase.email, client)

			if testCase.wantEmail == nil {
				client.assertNotCalled(suite.T())
			} else {
				client.assertCalledWithParams(suite.T(), *testCase.wantEmail)
			}
		})
	}
}

func TestGeneratePasswordReset(t *testing.T) {
	suite.Run(t, new(PasswordResetSuite))
}
