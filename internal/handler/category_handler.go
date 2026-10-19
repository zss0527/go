package handler

import (
	"bunny-fairy-api/internal/dto"
	"bunny-fairy-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CategoryHandler 处理器结构体，存放service【值类型】
type CategoryHandler struct {
	categoryService service.CategoryService
}

// NewCategoryHandler 构造函数，接收 service.CategoryService（值，不是指针）
func NewCategoryHandler(svc service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: svc,
	}
}

// GetCategoryTree 查询分类树接口
func (h *CategoryHandler) GetCategoryTree(c *gin.Context) {
	treeData, err := h.categoryService.GetCategoryTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "查询分类树失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取成功",
		"data": treeData,
	})
}

// ImportCategoryTree 导入树形分类数据接口
func (h *CategoryHandler) ImportCategoryTree(c *gin.Context) {
	var data []dto.CategoryDTO
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "JSON解析失败:" + err.Error(),
		})
		return
	}

	err := h.categoryService.ImportTree(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "导入失败:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "导入成功",
	})
}
