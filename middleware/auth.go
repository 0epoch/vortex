package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vortex/pkg/auth"
)

func Auth() gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if "" == tokenStr {
			tokenStr = c.Query("token")
		}
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : 401,
				"msg": "授权失败!",
			})
			c.Abort()
			return
		}
		tokenStr = tokenStr[7:]
		token, claims, err := auth.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg": "授权失败！",
			})
			c.Abort()
			return
		}
		userID := claims.UserID
		c.Set("userID", userID)
		c.Next()
	}
}