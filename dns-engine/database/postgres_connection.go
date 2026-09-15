package database

import (
	"database/sql"
)

type PostgreSQL struct {
	DB *sql.DB
}

func NewPostgreSQL(db *sql.DB) *PostgreSQL {
	return &PostgreSQL{DB: db}
}
