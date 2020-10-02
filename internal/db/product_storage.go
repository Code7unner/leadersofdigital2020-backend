package db

import (
	"database/sql"
	"errors"
)

const (
	insertProductQuery = `
		INSERT INTO products ("id", "name", "type", "description", "price", "img_url", "additional_info")
		VALUES($1, $2, $3, $4, $5, $6, $7);`
)

type ProductStorage struct {
	conn      *sql.DB
	tableName string
}

func (s *ProductStorage) Insert(row DBRow) error {
	dbRow, ok := row.(*Product)
	if !ok {
		return errors.New("product validation error")
	}

	bdTx, err := s.conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(
		insertProductQuery,
		dbRow.Id,
		dbRow.Name,
		dbRow.Type,
		dbRow.Description,
		dbRow.Price,
		dbRow.ImgUrl,
		dbRow.AdditionalInfo)
	if err != nil {
		return err
	}

	if err := bdTx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewProductStorage(conn *sql.DB) Storage {
	return &ProductStorage{
		conn:      conn,
		tableName: "products",
	}
}
