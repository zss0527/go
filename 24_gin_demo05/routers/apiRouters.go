package routers

import (
	"gindemo05/controllers/api"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

func ApiRouters(r *gin.Engine) {
	/*
		config middleware in router group
		1.
		piRouters := r.Group("/v1",middlewareFnHere){}
		2.
		yourRouters : = r.Group("/v1")
		yourRouters.Use(middleware)
	*/
	//apiRouters := r.Group("/v1",middlewareFnHere)
	apiRouters := r.Group("/v1")
	{
		//retrive params in get request
		apiRouters.GET("/getParams", api.ApiController{}.Params1)
		// func(ctx *gin.Context) {
		// 	userName := ctx.Query("username")
		// 	age := ctx.Query("age")
		// 	password := ctx.DefaultQuery("password", "123456")
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"userName": userName,
		// 		"age":      age,
		// 		"password": password,
		// 	})
		// })

		//retrive params in get request
		apiRouters.GET("/getParams1", api.ApiController{}.Params2)
		// func(ctx *gin.Context) {
		// 	var user UserModel
		// 	ctx.ShouldBind(&user)
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"userName": user.UserName,
		// 		"address":  user.Address,
		// 		"password": user.Password,
		// 	})
		// })
	}
}
