package middlewares

import (
	. "clokify/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authString := ctx.GetHeader("Authorization")
		token := strings.Split(authString, " ")
		if token[1] == "" {
			ctx.JSON(401, gin.H{
				"error": "Unauthorized. Missing token.",
			})
			ctx.Abort()
			return
		}

		_, isValidToken, err := ParseJWTToken(token[1])
		if err != nil {
			ctx.JSON(401, gin.H{
				"error": "Unauthorized. token mismatch",
			})
			return
		} else if isValidToken {
			ctx.Next()
		} else {
			ctx.JSON(401, gin.H{
				"error": "some issue with token",
			})
			return
		}
	}
}
