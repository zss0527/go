package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type AdminController struct {
	BaseController
}

func (con AdminController) Admin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
		"title": "index page",
	})
}

func (con AdminController) News(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "news page",
	})
}

func (con AdminController) AdminNews(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/news.html", gin.H{
		"title": "news page",
	})
}
