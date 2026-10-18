package models

import (
	"time"
)

// RolePermission 角色-权限关联表
type RolePermission struct {
	RoleID       uint      `gorm:"primaryKey" json:"roleId"`
	PermissionID uint      `gorm:"primaryKey" json:"permissionId"`
	CreatedAt    time.Time `json:"createdAt"`
}

// TableName 指定表名
func (RolePermission) TableName() string {
	return "role_permissions"
}
