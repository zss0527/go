package group3

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserModel struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

type Group3Controller struct{}

func (con Group3Controller) User(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.DefaultPostForm("password", "654321")
	ctx.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}

func (con Group3Controller) UserJson(ctx *gin.Context) {
	var user UserModel
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erroer": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"username": user.UserName,
		"password": user.Password,
		"address":  user.Address,
	})
}
