package service

import (
	"github.com/dcerverizzo/microservices-fc/goapi/internal/database"
	"github.com/dcerverizzo/microservices-fc/goapi/internal/entity"
)

type ProductService struct {
	ProductDB database.ProductDB
}

func NewProductService(db database.ProductDB) *ProductService {
	return &ProductService{
		ProductDB: db,
	}
}

func (s *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := s.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := s.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := s.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) CreateProduct(name, description, category_id, image_url string, price float64) (*entity.Product, error) {
	product := entity.NewProduct(name, description, category_id, image_url, price)
	_, err := s.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
