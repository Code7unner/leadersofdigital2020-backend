package db

import (
	"database/sql"
	"errors"
	sq "github.com/Masterminds/squirrel"
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

func (s *ProductStorage) GetAllProducts() (products []DBRow, err error) {
	sqlQ, _, err := sq.Select("id", "name", "type", "description", "price", "image_url", "additional_info").
		From(s.tableName).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.conn.Query(sqlQ, nil)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Type, &product.Description,
			&product.Price, &product.ImgUrl, &product.AdditionalInfo)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductStorage) SelectByType(productType string) (products []DBRow, err error) {

	sqlQ, _, err := sq.Select("id", "name", "type", "description", "price", "image_url", "additional_info").
		From(s.tableName).
		Where("type = ?", productType).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.conn.Query(sqlQ, productType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Type, &product.Description,
			&product.Price, &product.ImgUrl, &product.AdditionalInfo)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *ProductStorage) GetProductsByOrder(orderId int64) (products []DBRow, err error) {
	sqlQ, _, err := sq.Select("products.id", "products.name", "products.type", "products.description", "products.price", "products.image_url", "products.additional_info").
		Join("order_product ON order_product.product_id = products.id", nil).
		Where("order_product.order_id = ?", orderId).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.conn.Query(sqlQ, orderId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.Id, &product.Name, &product.Type, &product.Description,
			&product.Price, &product.ImgUrl, &product.AdditionalInfo)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func NewProductStorage(conn *sql.DB) Storage {
	return &ProductStorage{
		conn:      conn,
		tableName: "products",
	}
}
