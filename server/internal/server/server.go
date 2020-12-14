package server

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bensaufley/catalg/server/internal/graph"
	"github.com/bensaufley/catalg/server/internal/graph/generated"
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		DB: db,
	}}))

	http.Handle("/graphiql", playground.Handler("GraphQL playground", "/api"))
	http.Handle("/api", srv)

	log.Infof("Listening on :%s\n", opts.Port)
	if err := http.ListenAndServe(":"+opts.Port, nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
