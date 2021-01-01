package main

import (
	"github.com/bensaufley/catalg/server/internal/server"
	"github.com/bensaufley/catalg/server/internal/stubbables"
)

func main() {
	server.Serve(server.Opts{
		DBHost:         stubbables.GetEnvWithDefault("DATABASE_HOST", "localhost"),
		DBName:         stubbables.MustGetEnv("DATABASE_NAME"),
		DBPassword:     stubbables.MustGetEnv("DATABASE_PASSWORD"),
		DBPort:         stubbables.GetEnvWithDefault("DATABASE_PORT", "5432"),
		DBUser:         stubbables.MustGetEnv("DATABASE_USER"),
		MailerServer:   stubbables.MustGetEnv("MAILER_SERVER"),
		MailerUsername: stubbables.MustGetEnv("MAILER_USERNAME"),
		MailerPassword: stubbables.MustGetEnv("MAILER_PASSWORD"),
		Port:           stubbables.GetEnvWithDefault("PORT", "8080"),
	})
}
