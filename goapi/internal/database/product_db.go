package database

import (
	"database/sql"

	"github.com/dcerverizzo/microservices-fc/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (pd *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	row := pd.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	rows, err := pd.db.Query("SELECT * FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products (id, name, description, price, category_id, image_url) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Price, product.CategoryID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pd *ProductDB) UpdateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := pd.db.Exec("UPDATE products SET name = ?, description = ?, price = ?, category_id = ?, image_url = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.CategoryID, product.ImageURL, product.ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pd *ProductDB) DeleteProduct(id string) error {
	_, err := pd.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
