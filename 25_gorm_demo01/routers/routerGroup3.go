package routers

import (
	"gormdemo01/controllers/group3"

	"github.com/gin-gonic/gin"
)

func RouterGroup3(r *gin.Engine) {
	routerGroup3 := r.Group("/v2")
	{
		// retrive params in post request with form data
		routerGroup3.POST("/user", group3.Group3Controller{}.User)

		// retrive params in post request with body json data
		routerGroup3.POST("/userJson", group3.Group3Controller{}.UserJson)
	}
}
