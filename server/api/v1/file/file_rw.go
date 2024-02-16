package file

import (
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/common/request"
	"github.com/Allen9012/ServerManager/server/model/common/response"
	"github.com/Allen9012/ServerManager/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileRWApi struct {
}

var FRWService = service.ServiceGroupApp.FileServiceGroup.FileRWService

// ListFiles @Tags File
// @Summary List files
// @Description 获取文件列表
// @Accept json
// @Param request body request.FileOption true "request"
// @Success 200 {object} response.FileInfo
// @Security ApiKeyAuth
// @Router /files/search [post]
func (Frw *FileRWApi) ListFiles(c *gin.Context) {
	var req request.FileOption
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := FRWService.GetFileList(&req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
