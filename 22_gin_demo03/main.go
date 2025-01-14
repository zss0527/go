package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
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
