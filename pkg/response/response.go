package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Send(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code":code, "message":message})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 0,"msg": "success", "data": data})
}

func Failed(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"code": 4000, "message": message, "data": ""})
}

func NotFond(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"code":4004, "message": message})
}