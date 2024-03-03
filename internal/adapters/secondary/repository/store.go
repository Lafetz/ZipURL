package repository

import "database/sql"

type Store struct {
	db *sql.DB
}

func NewDb(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
