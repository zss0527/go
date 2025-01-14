package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Cotent string `json:"content`
}

func main() {
	r := gin.Default()
	//load html files
	r.LoadHTMLGlob("./templates/*")
	// fmt.Println(r)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello Gin!")
	})

	r.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"userName": "Larry",
			"password": 123456,
			"address":  "china",
		})
	})

	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"field1": "value1",
			"field2": 33333,
		})
	})

	a := Article{
		Title:  "this ia a title-jsonp",
		Desc:   "this is a desc",
		Cotent: "this is content",
	}
	r.GET("/json3", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &a)
	})

	r.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, &a)
	})

	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "hello gin, i am xml",
		})
	})

	r.GET("/newsPage", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "backend data",
		})
	})

	r.Run(":8081")
}
