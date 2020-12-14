package server

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/log"
	"github.com/bensaufley/catalg/server/internal/models"
)

type Opts struct {
	DBHost     string
	DBName     string
	DBPassword string
	DBPort     string
	DBUser     string
	Port       string
}

func Serve(opts Opts) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", opts.DBUser, opts.DBPassword, opts.DBHost, opts.DBPort, opts.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Infof("Listening on :%s\n", opts.Port)
	if err := http.ListenAndServe(":"+opts.Port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
