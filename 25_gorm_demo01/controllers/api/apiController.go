package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

type ApiController struct{}

func (con ApiController) Params1(ctx *gin.Context) {
	userName := ctx.Query("username")
	age := ctx.Query("age")
	password := ctx.DefaultQuery("password", "123456")
	ctx.JSON(http.StatusOK, gin.H{
		"userName": userName,
		"age":      age,
		"password": password,
	})
}

func (con ApiController) Params2(ctx *gin.Context) {
	var user UserModel
	ctx.ShouldBind(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"userName": user.UserName,
		"address":  user.Address,
		"password": user.Password,
	})
}
