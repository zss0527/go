package group44

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type Group4Controller struct{}

func (con Group4Controller) XmlPost(ctx *gin.Context) {
	b, _ := ctx.GetRawData()
	article := Article{}
	if err := xml.Unmarshal(b, &article); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, article)
}

func (con Group4Controller) UserUid(ctx *gin.Context) {
	uid := ctx.Param("uid")
	ctx.JSON(http.StatusOK, gin.H{
		"uid": uid,
	})
}
