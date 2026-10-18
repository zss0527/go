package middleware

import (
	"bunny-fairy-api/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "未提供Token"})
			ctx.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		claims, err := auth.ParseToken(tokenString, secret)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "无效Token"})
			ctx.Abort()
			return
		}

		// 将用户信息注入上下文
		ctx.Set("userID", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
