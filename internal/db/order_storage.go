package db

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const (
	insertOrderQuery = `
		INSERT INTO orders ("id", "courier_id", "status")
		VALUES($1, $2, $3);`
	deleteOrderQuery = `
		DELETE FROM orders WHERE id = $1;`
)

type OrderStorage struct {
	conn      *sql.DB
	tableName string
}

func (s *OrderStorage) Insert(row DBRow) error {
	dbRow, ok := row.(Order)
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

func (s *OrderStorage) SelectById(id int64) (order DBRow, err error) {
	sqlQ, _, err := sq.Select("id", "courier_id", "status").
		From(s.tableName).
		Where("id = ?", id).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := Order{}

	err = s.conn.QueryRow(sqlQ, id).Scan(&row.Id, &row.CourierId, &row.Status)
	if err != nil {
		return nil, sql.ErrNoRows
	}

	order = &row

	return order, nil
}

func (s *OrderStorage) SelectByCourier(courierId int64) (orders []DBRow, err error) {
	sqlQ, _, err := sq.Select("id", "courier_id", "status").
		From(s.tableName).
		Where("courier_id = ?", courierId).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.conn.Query(sqlQ, courierId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order := Order{}
		err = rows.Scan(&order.Id, order.Status, order.CourierId)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (s *OrderStorage) DeleteById(id int64) error {
	bdTx, err := s.conn.Begin()
	if err != nil {
		return err
	}

	_, err = bdTx.Exec(
		deleteOrderQuery,
		id)
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
