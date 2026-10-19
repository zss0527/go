package service

import (
	"bunny-fairy-api/internal/models/home"

	"gorm.io/gorm"
)

type HomeService struct {
	db              *gorm.DB
	categoryService CategoryService
}

func NewHomeService(db *gorm.DB, service CategoryService) *HomeService {
	return &HomeService{
		db:              db,
		categoryService: service,
	}
}

func (h *HomeService) GetHomeCategoryHead() ([]home.Category, error) {
	categoryTree, err := h.categoryService.GetCategoryTree()

	if err != nil {
		return nil, err
	}

	return categoryTree, nil
}
