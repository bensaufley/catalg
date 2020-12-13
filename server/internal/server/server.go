package server

import (
	"log"
	"net/http"

	"github.com/bensaufley/catalg/server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = db.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	port := opts.Port
	if port == "" {
		port = "8080"
		log.Println("PORT was not set. Defaulting to 8080")
	}

	log.Printf("Listening on :%s\n", port)
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
