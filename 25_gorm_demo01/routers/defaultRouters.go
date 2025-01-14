package routers

import (
	"fmt"
	"gormdemo01/controllers/admin"
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

		defaultRouter.GET("/news", admin.AdminController{}.News)

		defaultRouter.GET("/admin", admin.AdminController{}.Admin)

		defaultRouter.GET("/admin/news", admin.AdminController{}.AdminNews)
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
