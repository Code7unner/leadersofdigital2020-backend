package db

import (
	"database/sql"
	"errors"
)

const (
	insertOrderQuery = `
		INSERT INTO orders ("id", "courier_id", "status")
		VALUES($1, $2, $3);`
)

type OrderStorage struct {
	conn      *sql.DB
	tableName string
}

func (s *OrderStorage) Insert(row DBRow) error {
	dbRow, ok := row.(*Order)
	if !ok {
		return errors.New("order validation failed")
	}

	bdTx, err := s.conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(insertOrderQuery, dbRow.Id, dbRow.CourierId, dbRow.Status)
	if err != nil {
		return err
	}

	if err := bdTx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewOrderStorage(conn *sql.DB) Storage {
	return &OrderStorage{
		conn:      conn,
		tableName: "orders",
	}
}
