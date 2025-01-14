package routers

import (
	"fmt"
	"gindemo05/controllers/admin"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

func DefaultRouters(r *gin.Engine) {
	defaultRouter := r.Group("/")
	{
		defaultRouter.GET("/", initMiddleware1, admin.ArticleController{}.DefaultNoPath)
		// func(ctx *gin.Context) {
		// ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		// 	"title": "index page",
		// 	"score": 5,
		// 	"hobby": []string{"eat", "sleep", "code"},
		// 	"newsList": []interface{}{
		// 		&Article{
		// 			Title:   "title1111",
		// 			Content: "content1111",
		// 		},
		// 		&Article{
		// 			Title:   "title222",
		// 			Content: "content2222",
		// 		},
		// 	},
		// 	"testSlice": []string{},
		// 	"news": &Article{
		// 		Title:   "title3333",
		// 		Content: "content3333",
		// 	},
		// 	"date": 1629423555,
		// })
		// })

		defaultRouter.GET("/news", admin.AdminController{}.News)
		//func(ctx *gin.Context) {
		// ctx.HTML(http.StatusOK, "default/news.html", gin.H{
		// 	"title": "news page",
		// })
		// })

		defaultRouter.GET("/admin", admin.AdminController{}.Admin)
		// func(ctx *gin.Context) {
		// 	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		// 		"title": "index page",
		// 	})
		// })

		defaultRouter.GET("/admin/news", admin.AdminController{}.AdminNews)
		// func(ctx *gin.Context) {
		// 	ctx.HTML(http.StatusOK, "admin/news.html", gin.H{
		// 		"title": "news page",
		// 	})
		// })
	}
}

// this func will called after related router mathced and before handler called.
func initMiddleware1(c *gin.Context) {
	startTime := time.Now().UnixMilli()

	//nexting middleware funcs
	c.Next() //divide into two parts
	//c.Abort()  //abord other codes
	endTime := time.Now().UnixMilli()

	fmt.Println("this request took time:", endTime-startTime)
}
func initMiddleware2(c *gin.Context) {

}
