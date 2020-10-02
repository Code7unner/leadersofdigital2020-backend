package db

import "database/sql"

type UserStorage struct {
	conn      *sql.DB
	tableName string
}

func(s *UserStorage) Insert (row DBRow) error {
	return nil
}

func NewUserStorage(conn *sql.DB) Storage {
	return &UserStorage{
		conn:      conn,
		tableName: "users",
	}
}
