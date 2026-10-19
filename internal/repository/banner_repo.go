package repository

import (
	"bunny-fairy-api/internal/models/home"

	"gorm.io/gorm"
)

type BannerRepo struct {
	db *gorm.DB
}

func NewBannerRepo(db *gorm.DB) *BannerRepo {
	return &BannerRepo{db: db}
}

func (r *BannerRepo) FindByType(bannerType string) ([]home.Banner, error) {
	var banner []home.Banner

	err := r.db.Where("type = ?", bannerType).Find(&banner).Error
	if err != nil {
		return nil, err
	}
	return banner, nil
}
