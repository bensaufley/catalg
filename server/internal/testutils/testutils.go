package testutils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/migrations"
)

// TestWrapper is a type for middleware-style wrappers to
// be used in TestMain
type TestWrapper func(func() int) func() int

// PrepareDB is a TestWrapper sets up and tears
// down the test database
func PrepareDB(cb func() int) func() int {
	return func() int {
		port := os.Getenv("DATABASE_PORT")
		if port == "" {
			port = "5432"
		}
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		dbName := os.Getenv("DATABASE_NAME")
		if err := db.Transaction(func(txdb *gorm.DB) error {
			log.Debugf("Dropping db %s…", dbName)
			if tx := txdb.Raw(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)); tx.Error != nil {
				return tx.Error
			}
			if tx := txdb.Raw(fmt.Sprintf("CREATE DATABASE %s", dbName)); tx.Error != nil {
				return txdb.Error
			}
			log.Debugf("done dropping db %s.", dbName)
			log.Debugf("Migrating db %s…", dbName)
			if e := migrations.Perform(txdb); e != nil {
				return e
			}
			log.Debugf("done migrating db %s.", dbName)

			return nil
		}); err != nil {
			panic(err)
		}

		code := cb()

		log.Debugf("Dropping db %s…", dbName)
		if err := db.Transaction(func(txdb *gorm.DB) error {
			tx := txdb.Raw(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName))
			return tx.Error
		}); err != nil {
			panic(err)
		}
		log.Debugf("done dropping db %s.", dbName)

		return code
	}
}

// WrapTests receives a *testing.M and a list of wrappers
// to apply from innermost to outermost
func WrapTests(m *testing.M, wrappers ...TestWrapper) {
	cb := m.Run
	for _, wrapper := range wrappers {
		cb = wrapper(cb)
	}
	code := cb()
	os.Exit(code)
}

func AssertError(test *testing.T, want interface{}, got error) bool {
	if w, ok := want.(bool); ok {
		if w {
			return assert.Error(test, got)
		}
		return assert.NoError(test, got)
	}
	w := want.(error)
	if w == nil {
		return assert.NoError(test, got)
	}
	return assert.EqualError(test, got, w.Error())
}
