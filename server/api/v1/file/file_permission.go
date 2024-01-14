package file

import (
	"github.com/Allen9012/ServerManager/server/global"
    "github.com/Allen9012/ServerManager/server/model/file"
    "github.com/Allen9012/ServerManager/server/model/common/request"
    fileReq "github.com/Allen9012/ServerManager/server/model/file/request"
    "github.com/Allen9012/ServerManager/server/model/common/response"
    "github.com/Allen9012/ServerManager/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type FilePermissionApi struct {
}

var FPService = service.ServiceGroupApp.FileServiceGroup.FilePermissionService


// CreateFilePermission 创建描述文件权限的信息
// @Tags FilePermission
// @Summary 创建描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FilePermission true "创建描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /FP/createFilePermission [post]
func (FPApi *FilePermissionApi) CreateFilePermission(c *gin.Context) {
	var FP file.FilePermission
	err := c.ShouldBindJSON(&FP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := FPService.CreateFilePermission(&FP); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFilePermission 删除描述文件权限的信息
// @Tags FilePermission
// @Summary 删除描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FilePermission true "删除描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /FP/deleteFilePermission [delete]
func (FPApi *FilePermissionApi) DeleteFilePermission(c *gin.Context) {
	var FP file.FilePermission
	err := c.ShouldBindJSON(&FP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := FPService.DeleteFilePermission(FP); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFilePermissionByIds 批量删除描述文件权限的信息
// @Tags FilePermission
// @Summary 批量删除描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /FP/deleteFilePermissionByIds [delete]
func (FPApi *FilePermissionApi) DeleteFilePermissionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := FPService.DeleteFilePermissionByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFilePermission 更新描述文件权限的信息
// @Tags FilePermission
// @Summary 更新描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body file.FilePermission true "更新描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /FP/updateFilePermission [put]
func (FPApi *FilePermissionApi) UpdateFilePermission(c *gin.Context) {
	var FP file.FilePermission
	err := c.ShouldBindJSON(&FP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := FPService.UpdateFilePermission(FP); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFilePermission 用id查询描述文件权限的信息
// @Tags FilePermission
// @Summary 用id查询描述文件权限的信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query file.FilePermission true "用id查询描述文件权限的信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /FP/findFilePermission [get]
func (FPApi *FilePermissionApi) FindFilePermission(c *gin.Context) {
	var FP file.FilePermission
	err := c.ShouldBindQuery(&FP)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reFP, err := FPService.GetFilePermission(FP.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reFP": reFP}, c)
	}
}

// GetFilePermissionList 分页获取描述文件权限的信息列表
// @Tags FilePermission
// @Summary 分页获取描述文件权限的信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query fileReq.FilePermissionSearch true "分页获取描述文件权限的信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /FP/getFilePermissionList [get]
func (FPApi *FilePermissionApi) GetFilePermissionList(c *gin.Context) {
	var pageInfo fileReq.FilePermissionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := FPService.GetFilePermissionInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
