package entity

import "github.com/google/uuid"

type Category struct {
	ID   string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID:   uuid.New().String(),
		Name: name,
	}
}

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
	ImageURL    string
}

func NewProduct(name string, description string, CategoryID string, imageURL string, price float64) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  CategoryID,
		ImageURL:    imageURL,
	}
}