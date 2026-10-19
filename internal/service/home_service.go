package service

import (
	"bunny-fairy-api/internal/models"

	"gorm.io/gorm"
)

type HomeService struct {
	db *gorm.DB
}

func NewHomeService(db *gorm.DB) *HomeService {
	return &HomeService{
		db: db,
	}
}

func (h *HomeService) GetHomeCategoryHead() ([]models.CategoryHead, error) {
	categoryHeads := []models.CategoryHead{
		{ID: 1, Name: "居家"},
		{ID: 2, Name: "美食"},
		{ID: 3, Name: "服饰"},
		{ID: 4, Name: "母婴"},
		{ID: 5, Name: "个护"},
		{ID: 6, Name: "严选"},
		{ID: 7, Name: "数码"},
		{ID: 8, Name: "运动"},
		{ID: 9, Name: "杂项"},
	}

	return categoryHeads, nil
}
