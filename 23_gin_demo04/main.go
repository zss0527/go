package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

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
		"UnixToTime": UnixToTime,
		"PrintlnFn":  PrintlnFn,
	})
	//load template files
	r.LoadHTMLGlob("./templates/**/*")
	//create static web server
	r.Static("/static", "./static")

	//gin route
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "index page",
			"score": 5,
			"hobby": []string{"eat", "sleep", "code"},
			"newsList": []interface{}{
				&Article{
					Title:   "title1111",
					Content: "content1111",
				},
				&Article{
					Title:   "title222",
					Content: "content2222",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "title3333",
				Content: "content3333",
			},
			"date": 1629423555,
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "news page",
		})
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "index page",
		})
	})

	r.GET("/admin/news", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "news page",
		})
	})

	//retrive params in get request
	r.GET("/getParams", func(ctx *gin.Context) {
		userName := ctx.Query("username")
		age := ctx.Query("age")
		password := ctx.DefaultQuery("password", "123456")
		ctx.JSON(http.StatusOK, gin.H{
			"userName": userName,
			"age":      age,
			"password": password,
		})
	})

	//retrive params in get request
	r.GET("/getParams1", func(ctx *gin.Context) {
		var user UserModel
		ctx.ShouldBind(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"userName": user.UserName,
			"address":  user.Address,
			"password": user.Password,
		})
	})

	//retrive params in post request with form data
	r.POST("/user", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.DefaultPostForm("password", "654321")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//retrive params in post request with body json data
	r.POST("/userJson", func(ctx *gin.Context) {
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
	})

	r.POST("/xmlPost", func(ctx *gin.Context) {
		b, _ := ctx.GetRawData()
		article := Article{}
		if err := xml.Unmarshal(b, &article); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
		ctx.JSON(http.StatusOK, article)
	})

	//dynamic route
	r.GET("/user/:uid", func(ctx *gin.Context) {
		uid := ctx.Param("uid")
		ctx.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
	})
	r.Run(":8081")
}

// self definied template func
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func PrintlnFn(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}
