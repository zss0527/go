package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello %v!", "Gin")
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "I am a get request, mainly to get resources.")
	})

	r.POST("/add", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "I am a post request, mainly to add resources.")
	})

	r.PUT("/edit", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "this is an put request, mainly to edit resources.")
	})

	r.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "I am an delete request, mainly to delete resources.")
	})

	r.Run(":8081")
}
