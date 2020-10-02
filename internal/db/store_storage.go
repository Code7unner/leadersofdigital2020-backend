package db

import "database/sql"

type StoreStorage struct {
	conn      *sql.DB
	tableName string
}

func(s *StoreStorage) Insert (row DBRow) error {
	return nil
}

func NewStoreStorage(conn *sql.DB) Storage {
	return &StoreStorage{
		conn:      conn,
		tableName: "stores",
	}
}
