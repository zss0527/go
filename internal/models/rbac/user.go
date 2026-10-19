package rbac

import (
	"time"
)

type User struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username      string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Password      string     `gorm:"type:varchar(255);not null" json:"-"` // json:"-" 表示不返回密码
	Email         string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	IsActive      bool       `gorm:"type:tinyint(1);default:1" json:"isActive"`
	CreatedAt     *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"createdAt"`
	LastUpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"lastUpdatedAt"`
	Token         string     `gorm:"type:text" json:"token"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID            uint       `json:"id"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	CreatedAt     *time.Time `json:"createdAt"`
	LastUpdatedAt *time.Time `json:"lastUpdatedAt"`
	IsActive      bool       `json:"isActive"`
	Token         string     `json:"token"`
}

// 自定义表名（可选，不写的话默认就是 users）
func (User) TableName() string {
	return "users"
}
