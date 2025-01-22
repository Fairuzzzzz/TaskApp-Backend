package middleware

import "github.com/gin-gonic/gin"

func SetDefaultUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("user_id", uint(1))
		ctx.Next()
	}
}
