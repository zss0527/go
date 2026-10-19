package router

import (
	"bunny-fairy-api/internal/handler"
	"bunny-fairy-api/internal/middleware"
	"bunny-fairy-api/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRouter 设置路由
func SetupRouter(db *gorm.DB, secret string) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 你前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 关键放行Authorization
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化服务
	authService := service.NewAuthService(db, secret)
	permissionService := service.NewPermissionService(db)
	homeService := service.NewHomeService(db)

	// 初始化控制器（用于认证）
	authHandler := handler.NewAuthHandler(authService)

	// 初始化处理器（用于权限）
	permissionHandler := handler.NewPermissionHandler(permissionService)

	homeHandler := handler.NewHomeHandler(homeService)

	// API路由组
	api := r.Group("/api")
	{
		// 公开路由
		api.POST("/login", authHandler.Login)
		api.POST("/register", authHandler.Register)

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(secret))
		{

			home := protected.Group("/home")
			{
				home.GET("/category/head", homeHandler.HomeCategoryHead)
			}
			// 用户相关
			protected.GET("/profile", func(c *gin.Context) {
				userID := c.GetUint("userID")
				username := c.GetString("username")
				c.JSON(200, gin.H{
					"userID":   userID,
					"username": username,
				})
			})

			// 权限相关路由
			// 获取所有权限
			protected.GET("/permissions", permissionHandler.GetAllPermissions)

			// 获取当前用户的权限
			protected.GET("/my-permissions", permissionHandler.GetMyPermissions)
			protected.GET("/my-permissions/details", permissionHandler.GetMyPermissionsWithDetails)

			// 根据用户ID获取权限（管理员功能）
			protected.GET("/users/:userId/permissions", permissionHandler.GetPermissionsByUserID)
			protected.GET("/users/:userId/permissions/details", permissionHandler.GetUserPermissionsWithDetails)
		}
	}

	return r
}
