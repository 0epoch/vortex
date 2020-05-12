package router

import (
	"github.com/gin-gonic/gin"
	"vortex/handle"
	"vortex/middleware"
)

const apiRoot = "/api/v1"

func Load() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	user := handle.UserHandle{}
	v1 := r.Group(apiRoot)
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Create)
	}

	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.POST("/user/group", user.Join)
	}

	//公共API
	common := handle.Common{}
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.GET("/ws", common.Ws)
		v1.POST("/upload", common.Upload)
	}

	group := handle.GroupHandle{}
	v1 = r.Group(apiRoot, middleware.Auth())
	{
		v1.POST("/group", group.Create)
	}
	return r
}

