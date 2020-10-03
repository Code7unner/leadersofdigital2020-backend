package db

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
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
func (s *UserStorage) GetUserById(id int64) (user DBRow, err error) {
	sqlQ, _, err := sq.Select("id", "name", "phone", "password", "address", "sex", "role").
		From(s.tableName).
		Where("id = ?", id).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := User{}

	err = s.conn.QueryRow(sqlQ, id).Scan(&row.Id, &row.Name, &row.Phone, &row.Password, &row.Address, &row.Sex, &row.Role)
	if err != nil {
		return nil, err
	}

	user = &row

	return user, err
}

func NewUserStorage(conn *sql.DB) Storage {
	return &UserStorage{
		conn:      conn,
		tableName: "users",
	}
}
