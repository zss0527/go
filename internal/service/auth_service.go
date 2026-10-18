package service

import (
	"errors"
	"time"

	"bunny-fairy-api/internal/auth"
	"bunny-fairy-api/internal/dto"
	"bunny-fairy-api/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService 认证服务
type AuthService struct {
	db     *gorm.DB
	secret string
}

// NewAuthService 创建认证服务实例
func NewAuthService(db *gorm.DB, secret string) *AuthService {
	return &AuthService{
		db:     db,
		secret: secret,
	}
}

func (s *AuthService) Register(username, password, email string) error {
	// 检查用户名是否已存在
	var user models.User
	err := s.db.Where("username = ?", username).First(&user).Error
	if err == nil {
		return errors.New("用户名已存在")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// 创建新用户
	now := time.Now()
	newUser := models.User{
		Username:      username,
		Password:      string(hash),
		Email:         email,
		IsActive:      true,
		CreatedAt:     &now,
		LastUpdatedAt: &now,
	}
	err = s.db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}

// Login 处理登录逻辑
func (s *AuthService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. 查询用户
	var user models.User
	err := s.db.Where("username = ? AND is_active = true", req.Username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 2. 验证密码
	//if user.Password != req.Password {
	//	return nil, errors.New("用户名或密码错误")
	//}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 3. 查询用户的角色列表
	var roles []string
	s.db.Table("roles").
		Joins("JOIN user_roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", user.ID).
		Pluck("roles.name", &roles)

	// 4. 查询用户的权限列表（去重）
	var permissions []string
	s.db.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", user.ID).
		Distinct().
		Pluck("permissions.name", &permissions)

	// 5. 生成JWT Token
	token, expiresAt, err := auth.GenerateToken(user.ID, user.Username, s.secret)
	if err != nil {
		return nil, err
	}

	// 6. 更新用户表的 token 和 last_updated_at 字段
	now := time.Now()
	err = s.db.Model(&user).Updates(map[string]interface{}{
		"token":           token,
		"last_updated_at": now,
	}).Error
	if err != nil {
		return nil, errors.New("更新用户token失败")
	}

	// 7. 构建响应
	response := &dto.LoginResponse{
		Token:     token,
		TokenType: "Bearer",
		ExpiresAt: expiresAt,
		User: dto.UserInfo{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			IsActive: user.IsActive,
			Roles:    roles,
		},
		Permissions: permissions,
	}

	return response, nil
}

// GetUserByID 根据ID获取用户
func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

