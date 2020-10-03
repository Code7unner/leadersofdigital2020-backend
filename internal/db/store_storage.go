package db

import (
	"database/sql"
	"errors"
)

const (
	insertStoreQuery = `
		INSERT INTO stores ("id", "name", "address")
		VALUES($1, $2, $3);`
)

type StoreStorage struct {
	conn      *sql.DB
	tableName string
}

func (s *StoreStorage) Insert(row DBRow) error {
	dbRow, ok := row.(Store)
	if !ok {
		return errors.New("store validation error")
	}

	bdTx, err := s.conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(
		insertStoreQuery,
		dbRow.Id,
		dbRow.Name,
		dbRow.Address)
	if err != nil {
		return err
	}

	if err := bdTx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewStoreStorage(conn *sql.DB) Storage {
	return &StoreStorage{
		conn:      conn,
		tableName: "stores",
	}
}
