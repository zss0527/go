package service

import (
	"bunny-fairy-api/internal/models/rbac"
	"bunny-fairy-api/internal/repository"
	"errors"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Authenticate(req rbac.LoginRequest) (*rbac.UserResponse, error) {
	user, err := s.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// 实际项目中密码应为 bcrypt 比较
	if user.Password != req.Password {
		return nil, errors.New("invalid username or password")
	}

	return &rbac.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		LastUpdatedAt: user.LastUpdatedAt,
		IsActive:      user.IsActive,
		Token:         user.Token,
	}, nil
}

func (s *UserService) GetByID(id uint) (*rbac.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &rbac.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		LastUpdatedAt: user.LastUpdatedAt,
		IsActive:      user.IsActive,
		Token:         user.Token,
	}, nil
}

func (s *UserService) UpdateLoginTime(id uint, loginTime time.Time) (*rbac.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.LastUpdatedAt = &loginTime
	if err := s.repo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	user, err = s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &rbac.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		LastUpdatedAt: user.LastUpdatedAt,
		IsActive:      user.IsActive,
		Token:         user.Token,
	}, nil
}

func (s *UserService) Update(id uint, email string) (*rbac.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Email = email
	if err := s.repo.Update(user); err != nil {
		return nil, errors.New("failed to update user")
	}

	return &rbac.UserResponse{
		ID:            user.ID,
		Username:      user.Username,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		LastUpdatedAt: user.LastUpdatedAt,
		IsActive:      user.IsActive,
		Token:         user.Token,
	}, nil
}
