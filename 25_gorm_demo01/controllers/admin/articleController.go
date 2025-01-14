package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) DefaultNoPath(ctx *gin.Context) {
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
}

func (con ArticleController) Test(ctx *gin.Context) {
	var article Article

	if err := ctx.ShouldBind(&article); err != nil {
		con.failed(ctx)
	}
	con.success(ctx)
}
