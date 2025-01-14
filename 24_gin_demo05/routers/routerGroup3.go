package routers

import (
	"gindemo05/controllers/group3"

	"github.com/gin-gonic/gin"
)

func RouterGroup3(r *gin.Engine) {
	routerGroup3 := r.Group("/v2")
	{
		// retrive params in post request with form data
		routerGroup3.POST("/user", group3.Group3Controller{}.User)
		// func(ctx *gin.Context) {
		// 	username := ctx.PostForm("username")
		// 	password := ctx.DefaultPostForm("password", "654321")
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"username": username,
		// 		"password": password,
		// 	})
		// })

		// retrive params in post request with body json data
		routerGroup3.POST("/userJson", group3.Group3Controller{}.UserJson)
		// func(ctx *gin.Context) {
		// 	var user UserModel
		// 	if err := ctx.ShouldBindJSON(&user); err != nil {
		// 		ctx.JSON(http.StatusBadRequest, gin.H{
		// 			"erroer": err.Error(),
		// 		})
		// 		return
		// 	}
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"username": user.UserName,
		// 		"password": user.Password,
		// 		"address":  user.Address,
		// 	})
		// })
	}
}
