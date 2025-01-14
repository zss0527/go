package routers

import (
	"gormdemo01/controllers/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouters(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		adminRouters.GET("/users", admin.UserController{}.Users)
		adminRouters.DELETE("/user/:uid", admin.UserController{}.DeleteUser)
		adminRouters.POST("/user", admin.UserController{}.AddUser)
		adminRouters.PUT("/user", admin.UserController{}.EditUser)
	}
}
