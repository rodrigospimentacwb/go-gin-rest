package repository

import (
	"database/sql"
	"fmt"
	"github.com/rodrigospimentacwb/go-gin-rest/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return []model.Product{}, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var productList []model.Product
	for rows.Next() {
		var p model.Product
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price)
		if err != nil {
			return []model.Product{}, fmt.Errorf("failed to scan row: %w", err)
		}
		productList = append(productList, p)
	}

	if err = rows.Err(); err != nil {
		return []model.Product{}, fmt.Errorf("rows iteration error: %w", err)
	}

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := pr.connection.Prepare(
		"INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare(
		"SELECT id, product_name, price FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var product model.Product
	err = query.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	query.Close()
	return &product, nil
}
