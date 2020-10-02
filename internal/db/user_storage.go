package db

import (
	"database/sql"
	"errors"
)

const (
	insertUserQuery = `
		INSERT INTO users ("id", "name", "phone", "password", "address", "sex", "role")
		VALUES($1, $2, $3, $4, $5, $6, $7);`
)

type UserStorage struct {
	conn      *sql.DB
	tableName string
}

func (s *UserStorage) Insert(row DBRow) error {
	dbRow, ok := row.(*User)
	if !ok {
		return errors.New("user validation error")
	}

	bdTx, err := s.conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(
		insertUserQuery,
		dbRow.Id,
		dbRow.Name,
		dbRow.Phone,
		dbRow.Password,
		dbRow.Address,
		dbRow.Sex,
		dbRow.Role)
	if err != nil {
		return err
	}

	if err := bdTx.Commit(); err != nil {
		return err
	}

	return nil
}

func NewUserStorage(conn *sql.DB) Storage {
	return &UserStorage{
		conn:      conn,
		tableName: "users",
	}
}
