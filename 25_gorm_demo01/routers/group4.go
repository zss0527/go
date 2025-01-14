package routers

import (
	group44 "gormdemo01/controllers/group4"

	"github.com/gin-gonic/gin"
)

func Group4(r *gin.Engine) {
	group4 := r.Group("/v3")
	{
		group4.POST("/xmlPost", group44.Group4Controller{}.XmlPost)

		//dynamic route
		group4.GET("/user/:uid", group44.Group4Controller{}.UserUid)
	}
}
