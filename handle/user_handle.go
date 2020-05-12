package handle

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"vortex/model"
	"vortex/pkg/auth"
	"vortex/pkg/response"
	"vortex/service"
)

type UserHandle struct {}

func (uh *UserHandle) Create(c *gin.Context) {
	userBind := &model.User{}
	err := c.ShouldBind(userBind)
	if err != nil {
		response.Failed(c, "添加失败！")
		return
	}

	data := map[string]interface{}{
		"username": userBind.Username,
		"password": userBind.Password,
	}
	user, err := service.CreateUser(data)
	if err != nil {
		response.Failed(c, "注册失败！")
		return
	}
	response.Success(c, user)
}

//登录
func (uh *UserHandle) Login(c *gin.Context) {
	pwd := c.PostForm("password")
	user := &model.User{}
	err := model.DB.Where("username = ?", c.PostForm("username")).First(user).Error
	if err != nil {
		response.NotFond(c,"用户不存在！")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd))
	if err != nil {
		response.Send(c, 4000,"密码错误", "")
		return
	}

	token, er := auth.GenerateToken(user)
	if er != nil {
		response.Send(c, 4001,"授权失败！", "")
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user": user,
	}
	response.Success(c, data)
}

func (uh *UserHandle) Join(c *gin.Context) {
	bind := &model.UserGroup{}
	err := c.ShouldBind(bind)
	if err != nil {
		response.Failed(c, "加入失败！")
		return
	}

	data := map[string]interface{}{
		"userID": bind.UserID,
		"groupID": bind.GroupID,
	}

	err = service.JoinGroup(data)
	if err != nil {
		response.Failed(c, "加入失败！")
		return
	}
	response.Success(c, "")
}