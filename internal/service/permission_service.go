package service

import (
	"bunny-fairy-api/internal/models/rbac"

	"gorm.io/gorm"
)

// PermissionService 权限服务
type PermissionService struct {
	db *gorm.DB
}

// NewPermissionService 创建权限服务实例
func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{
		db: db,
	}
}

// GetAllPermissions 获取数据库里面所有的permission
func (s *PermissionService) GetAllPermissions() ([]rbac.Permission, error) {
	var permissions []rbac.Permission

	err := s.db.Order("id").Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// GetPermissionsByUserID 根据用户的id获取对应的permission
func (s *PermissionService) GetPermissionsByUserID(userID uint) ([]string, error) {
	var permissions []string

	err := s.db.Table("permissions").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Distinct().
		Pluck("permissions.name", &permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}

// GetUserPermissionsWithDetails 根据用户的id获取permission的详细信息（包含id, name, resource, action）
func (s *PermissionService) GetUserPermissionsWithDetails(userID uint) ([]rbac.Permission, error) {
	var permissions []rbac.Permission

	err := s.db.Table("permissions").
		Select("permissions.*").
		Joins("JOIN role_permissions ON permissions.id = role_permissions.permission_id").
		Joins("JOIN user_roles ON role_permissions.role_id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Distinct().
		Order("permissions.id").
		Find(&permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
