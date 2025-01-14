package router

import (
	"github.com/gin-gonic/gin"
	// "gin_gorm_demo/internal/middleware"
	"gin_gorm_demo/internal/controller"
)

func RegisterUserRoutes(r *gin.Engine, userCtrl *controller.UserController) {
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", userCtrl.Register)
		userGroup.POST("/login", userCtrl.Login)
		userGroup.POST("/add", userCtrl.AddUser)
		// 示例：添加需要认证的路由
		// authGroup := userGroup.Group("/protected", middleware.AuthMiddleware())
		// {
		//     // 这里可以添加需要认证的用户相关路由
		// }
	}
}
