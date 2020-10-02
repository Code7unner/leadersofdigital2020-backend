package db

import "database/sql"

type ProductStorage struct {
	conn      *sql.DB
	tableName string
}

func(s *ProductStorage) Insert (row DBRow) error {
	return nil
}

func NewProductStorage(conn *sql.DB) Storage {
	return &ProductStorage{
		conn:      conn,
		tableName: "products",
	}
}
