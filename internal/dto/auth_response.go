package dto

// LoginResponse 登录响应
type LoginResponse struct {
	Token       string   `json:"token"`
	TokenType   string   `json:"tokenType"`
	ExpiresAt   int64    `json:"expiresAt"`
	User        UserInfo `json:"user"`
	Permissions []string `json:"permissions"`
}

// UserInfo 用户信息（嵌套在LoginResponse中）
type UserInfo struct {
	ID       uint     `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	IsActive bool     `json:"isActive"`
	Roles    []string `json:"roles"`
}
