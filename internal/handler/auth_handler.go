package handler

import (
	"net/http"

	"bunny-fairy-api/internal/dto"
	"bunny-fairy-api/internal/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证控制器
type AuthHandler struct {
	authService *service.AuthService
}

// NewAuthHandler 创建认证控制器实例
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Login 登录接口
// POST /api/login
func (c *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.authService.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    resp,
	})
}

func (c *AuthHandler) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.authService.Register(req.Username, req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
