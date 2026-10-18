package models

import (
	"time"
)

// Permission 权限表
type Permission struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"unique;not null" json:"name"` // 格式：livestock:view
	Resource    string    `json:"resource"`                    // 资源：livestock, environment, device
	Action      string    `json:"action"`                      // 操作：view, create, update, delete
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName 指定表名
func (Permission) TableName() string {
	return "permissions"
}
