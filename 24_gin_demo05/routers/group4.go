package routers

import (
	group44 "gindemo05/controllers/group4"

	"github.com/gin-gonic/gin"
)

func Group4(r *gin.Engine) {
	group4 := r.Group("/v3")
	{
		group4.POST("/xmlPost", group44.Group4Controller{}.XmlPost)
		// func(ctx *gin.Context) {
		// 	b, _ := ctx.GetRawData()
		// 	article := Article{}
		// 	if err := xml.Unmarshal(b, &article); err != nil {
		// 		ctx.JSON(http.StatusBadRequest, err.Error())
		// 	}
		// 	ctx.JSON(http.StatusOK, article)
		// })

		//dynamic route
		group4.GET("/user/:uid", group44.Group4Controller{}.UserUid)
		// func(ctx *gin.Context) {
		// 	uid := ctx.Param("uid")
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"uid": uid,
		// 	})
		// })
	}
}
