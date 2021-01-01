package testutils

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/khaiql/dbcleaner.v2/engine"
)

// Postgres dbEngine
type Postgres struct {
	db *sql.DB
}

// NewPostgresEngine returns engine for Postgres db
func NewPostgresEngine(dsn string) engine.Engine {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Truncate(table string) error {
	cmd := fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table)

	_, err := p.db.Exec(cmd)
	return err
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
