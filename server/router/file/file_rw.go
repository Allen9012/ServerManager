package file

import (
	v1 "github.com/Allen9012/ServerManager/server/api/v1"
	"github.com/Allen9012/ServerManager/server/middleware"
	"github.com/gin-gonic/gin"
)

type FileRWRouter struct {
}

func (s *FileRWRouter) InitFileRWRouter(Router *gin.RouterGroup) {
	FileRouter := Router.Group("files").Use(middleware.OperationRecord())
	FileRouterWithoutRecord := Router.Group("files")
	var fileApi = v1.ApiGroupApp.FileApiGroup.FileRWApi
	{
		FileRouter.POST("create", fileApi.CreateFile)
	}
	{
		FileRouterWithoutRecord.POST("search", fileApi.ListFiles) // 获取文件列表
	}
}
