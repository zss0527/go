package handler

import (
	"bunny-fairy-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BannerHandler struct {
	bannerService *service.BannerService
}

func NewBannerHandler(bannerService *service.BannerService) *BannerHandler {
	return &BannerHandler{
		bannerService: bannerService,
	}
}

func (h *BannerHandler) FindBannerByType(c *gin.Context) {
	bannerType := c.Query("distributionSite")

	banner, err := h.bannerService.FindByType(bannerType)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "search failed " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "search success",
		"data": banner,
	})
}
