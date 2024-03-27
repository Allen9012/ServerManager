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
		FileRouter.POST("", fileApi.CreateFile)                         // 创建文件
		FileRouter.POST("/del", fileApi.DeleteFile)                     // 删除文件
		FileRouter.POST("/batch/del", fileApi.BatchDeleteFile)          // 批量删除文件
		FileRouter.POST("/move", fileApi.MoveFile)                      // 移动文件
		FileRouter.POST("/rename", fileApi.ChangeFileName)              // 复制文件
		FileRouter.POST("/upload", fileApi.UploadFiles)                 // 上传文件
		FileRouter.POST("/owner", fileApi.ChangeFileOwner)              // 修改权限
		FileRouter.POST("/mode", fileApi.ChangeFileMode)                // 修改文件权限
		FileRouter.POST("/batch/role", fileApi.BatchChangeModeAndOwner) // 批量修改文件权限和用户/组
	}
	{
		FileRouterWithoutRecord.POST("/search", fileApi.ListFiles)   // 获取文件列表
		FileRouterWithoutRecord.POST("/check", fileApi.CheckFile)    // 检查文件是否存在
		FileRouterWithoutRecord.POST("/content", fileApi.GetContent) // 获取文件属性
		FileRouterWithoutRecord.GET("/download", fileApi.Download)   // 下载文件
	}
}
