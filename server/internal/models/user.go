package models

import "database/sql"

type User struct {
	Model
	Username       string
	Email          string
	PasswordDigest sql.NullString
	Pepper         sql.NullString
	ActivatedAt    sql.NullTime
}
