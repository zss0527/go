package handler

import (
	"bunny-fairy-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	homeService *service.HomeService
}

func NewHomeHandler(homeService *service.HomeService) *HomeHandler {
	return &HomeHandler{
		homeService: homeService,
	}
}

func (h *HomeHandler) HomeCategoryHead(c *gin.Context) {
	categoryHeads, err := h.homeService.GetHomeCategoryHead()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取分类头部列表失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data":    categoryHeads,
	})
}
