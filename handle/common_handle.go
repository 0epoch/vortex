package handle

import (
	"github.com/gin-gonic/gin"
	"vortex/im"
	"vortex/pkg/response"
	"vortex/pkg/vortexfile"
)

type Common struct {}

func(ch *Common) Ws(c *gin.Context) {
	im.ServeWs(c)
}

func (ch *Common) Upload(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	err := vortexfile.SizeLimit(fileHeader, "image")
	if err != nil {
		response.Failed(c, err.Error())
		return
	}
	filePath := vortexfile.FullPath(fileHeader)
	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		response.Failed(c, "上传失败！")
		return
	}
	response.Success(c, "/"+filePath)
}