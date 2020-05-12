package handle

import (
	"github.com/gin-gonic/gin"
	"vortex/model"
	"vortex/pkg/response"
	"vortex/service"
)

type GroupHandle struct {}

func (g *GroupHandle) Create(c *gin.Context) {
	groupBind := &model.Group{}
	err := c.ShouldBind(groupBind)

	if err != nil {
		response.Failed(c, "群创建失败！")
		return
	}
	userID, _ := c.Get("userID")
	groupBind.CreatedUid = userID.(int64)
	data := map[string]interface{}{
		"name": groupBind.Name,
		"createdUid": groupBind.CreatedUid,
	}
	group, err := service.CreateGroup(data)
	if err != nil {
		response.Failed(c, "群创建失败！")
		return
	}
	response.Success(c, group)
}

