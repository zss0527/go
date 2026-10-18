package handler

import "github.com/gin-gonic/gin"

// Handler 接口：每个业务 handler 都要实现
type Handler interface {
	RegisterRoutes(r *gin.RouterGroup)
}

// RegisterAll 统一注册所有 handler
func RegisterAll(r *gin.Engine, handlers ...Handler) {
	api := r.Group("/api")
	for _, h := range handlers {
		h.RegisterRoutes(api)
	}
}
