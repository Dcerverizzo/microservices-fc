package service

import (
	"github.com/dcerverizzo/microservices-fc/goapi/internal/database"
	"github.com/dcerverizzo/microservices-fc/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(db database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: db,
	}
}

func (s *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := s.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)
	_, err := s.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := s.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
