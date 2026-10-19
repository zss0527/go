package service

import (
	"bunny-fairy-api/internal/models/home"
	"bunny-fairy-api/internal/repository"
)

type BannerService struct {
	bannerRepo *repository.BannerRepo
}

func NewBannerService(bannerRepo *repository.BannerRepo) *BannerService {
	return &BannerService{bannerRepo: bannerRepo}
}

func (s *BannerService) FindByType(bannerType string) ([]home.Banner, error) {
	bannerList, err := s.bannerRepo.FindByType(bannerType)

	if err != nil {
		return nil, err
	}

	return bannerList, nil
}
