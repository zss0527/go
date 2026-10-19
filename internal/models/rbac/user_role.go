package rbac

import (
	"time"
)

// UserRole 用户-角色关联表
type UserRole struct {
	UserID    uint      `gorm:"primaryKey" json:"userId"`
	RoleID    uint      `gorm:"primaryKey" json:"roleId"`
	GrantedAt time.Time `json:"grantedAt"`
}

// TableName 指定表名
func (UserRole) TableName() string {
	return "user_roles"
}
