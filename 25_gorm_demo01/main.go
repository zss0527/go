package main

import (
	"gormdemo01/models"
	"gormdemo01/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type UserModel struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Address  string `form:"address" json:"address"`
}

func main() {

	r := gin.Default()
	//self definied template funcs must before r.LoadHTML()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"PrintlnFn":  models.PrintlnFn,
	})
	//load template files
	r.LoadHTMLGlob("./templates/**/*")
	//create static web server
	r.Static("/static", "./static")

	//global middlware
	r.Use(models.GlobalMiddleWare1)

	routers.AdminRouters(r)
	routers.DefaultRouters(r)
	routers.ApiRouters(r)

	routers.RouterGroup3(r)
	routers.Group4(r)
	r.Run(":8081")
}
