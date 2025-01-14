package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (con BaseController) success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}

func (con BaseController) failed(c *gin.Context) {
	c.String(http.StatusInternalServerError, "failed")
}
