package handler

import (
	"net/http"
	"strconv"

	"bunny-fairy-api/internal/service"

	"github.com/gin-gonic/gin"
)

// PermissionHandler 权限处理器
type PermissionHandler struct {
	permissionService *service.PermissionService
}

// NewPermissionHandler 创建权限处理器实例
func NewPermissionHandler(permissionService *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

// GetAllPermissions 获取所有权限
// GET /api/permissions
func (h *PermissionHandler) GetAllPermissions(c *gin.Context) {
	permissions, err := h.permissionService.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取权限列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    permissions,
	})
}

// GetPermissionsByUserID 根据用户ID获取权限
// GET /api/users/:userId/permissions
func (h *PermissionHandler) GetPermissionsByUserID(c *gin.Context) {
	// 获取路径参数
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
			"error":   err.Error(),
		})
		return
	}

	// 调用service获取权限
	permissions, err := h.permissionService.GetPermissionsByUserID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户权限失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    permissions,
	})
}

// GetUserPermissionsWithDetails 根据用户ID获取权限详细信息
// GET /api/users/:userId/permissions/details
func (h *PermissionHandler) GetUserPermissionsWithDetails(c *gin.Context) {
	// 获取路径参数
	userIDStr := c.Param("userId")
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的用户ID",
			"error":   err.Error(),
		})
		return
	}

	// 调用service获取权限详细信息
	permissions, err := h.permissionService.GetUserPermissionsWithDetails(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户权限详情失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    permissions,
	})
}

// GetMyPermissions 获取当前登录用户的权限（从token中获取userID）
// GET /api/my-permissions
func (h *PermissionHandler) GetMyPermissions(c *gin.Context) {
	// 从中间件设置的上下文中获取userID
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	// 类型断言
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "用户ID类型错误",
		})
		return
	}

	// 调用service获取权限
	permissions, err := h.permissionService.GetPermissionsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户权限失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    permissions,
	})
}

// GetMyPermissionsWithDetails 获取当前登录用户的权限详细信息
// GET /api/my-permissions/details
func (h *PermissionHandler) GetMyPermissionsWithDetails(c *gin.Context) {
	// 从中间件设置的上下文中获取userID
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "未登录或登录已过期",
		})
		return
	}

	// 类型断言
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "用户ID类型错误",
		})
		return
	}

	// 调用service获取权限详细信息
	permissions, err := h.permissionService.GetUserPermissionsWithDetails(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取用户权限详情失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    permissions,
	})
}
