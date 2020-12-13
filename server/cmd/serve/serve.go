package main

import (
	"os"

	"github.com/bensaufley/catalg/server/internal/server"
)

func main() {
	server.Serve(server.Opts{
		DBHost:     os.Getenv("DATABASE_HOST"),
		DBName:     os.Getenv("DATABASE_NAME"),
		DBPassword: os.Getenv("DATABASE_PASSWORD"),
		DBPort:     os.Getenv("DATABASE_PORT"),
		DBUser:     os.Getenv("DATABASE_USER"),
		Port:       os.Getenv("PORT"),
	})
}
