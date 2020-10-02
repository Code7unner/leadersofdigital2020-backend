package db

import "database/sql"

type OrderStorage struct {
	conn      *sql.DB
	tableName string
}

func(s *OrderStorage) Insert (row DBRow) error {
	return nil
}

func NewOrderStorage(conn *sql.DB) Storage {
	return &OrderStorage{
		conn:      conn,
		tableName: "orders",
	}
}
