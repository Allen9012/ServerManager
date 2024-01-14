package file

import (
	"github.com/Allen9012/ServerManager/server/api/v1"
	"github.com/Allen9012/ServerManager/server/middleware"
	"github.com/gin-gonic/gin"
)

type FilePermissionRouter struct {
}

// InitFilePermissionRouter 初始化 描述文件权限的信息 路由信息
func (s *FilePermissionRouter) InitFilePermissionRouter(Router *gin.RouterGroup) {
	FPRouter := Router.Group("FP").Use(middleware.OperationRecord())
	FPRouterWithoutRecord := Router.Group("FP")
	var FPApi = v1.ApiGroupApp.FileApiGroup.FilePermissionApi
	{
		FPRouter.POST("createFilePermission", FPApi.CreateFilePermission)   // 新建描述文件权限的信息
		FPRouter.DELETE("deleteFilePermission", FPApi.DeleteFilePermission) // 删除描述文件权限的信息
		FPRouter.DELETE("deleteFilePermissionByIds", FPApi.DeleteFilePermissionByIds) // 批量删除描述文件权限的信息
		FPRouter.PUT("updateFilePermission", FPApi.UpdateFilePermission)    // 更新描述文件权限的信息
	}
	{
		FPRouterWithoutRecord.GET("findFilePermission", FPApi.FindFilePermission)        // 根据ID获取描述文件权限的信息
		FPRouterWithoutRecord.GET("getFilePermissionList", FPApi.GetFilePermissionList)  // 获取描述文件权限的信息列表
	}
}
