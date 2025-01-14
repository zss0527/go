package routers

import (
	"gormdemo01/controllers/api"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

func ApiRouters(r *gin.Engine) {
	apiRouters := r.Group("/v1")
	{
		apiRouters.GET("/getParams", api.ApiController{}.Params1)
		apiRouters.GET("/getParams1", api.ApiController{}.Params2)
	}
}
