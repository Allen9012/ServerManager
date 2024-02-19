package file

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Allen9012/ServerManager/server/buserr"
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
	files, err := FRWService.GetFileList(req)
	if err != nil {
		global.GVA_LOG.Error("ListFiles fail", zap.Error(err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(files, c)
}

// @Tags File
// @Summary Create file
// @Description 创建文件/文件夹
// @Accept json
// @Param request body request.FileCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"创建文件/文件夹 [path]","formatEN":"Create dir or file [path]"}
func (Frw *FileRWApi) CreateFile(c *gin.Context) {
	var req request.FileCreate
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FRWService.Create(req)
	if err != nil {
		global.GVA_LOG.Error("CreateFile fail", zap.Error(err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// @Tags File
// @Summary Delete file
// @Description 删除文件/文件夹
// @Accept json
// @Param request body request.FileDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/del [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"删除文件/文件夹 [path]","formatEN":"Delete dir or file [path]"}
// func (Frw *FileRWApi) DeleteFile(c *gin.Context) {
// var req request.FileDelete
// if err := helper.CheckBindAndValidate(&req, c); err != nil {
// 	return
// }
// err := fileService.Delete(req)
// if err != nil {
// 	helper.ErrorWithDetail(c, constant.CodeErrInternalServer, constant.ErrTypeInternalServer, err)
// 	return
// }
// helper.SuccessWithData(c, nil)
// }

// TODO modify logic
// @Tags File
// @Summary Upload file
// @Description 上传文件
// @Param file formData file true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/upload [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"上传文件 [path]","formatEN":"Upload file [path]"}
func (Frw *FileRWApi) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	files := form.File["file"]
	paths := form.Value["path"]
	if len(paths) == 0 || !strings.Contains(paths[0], "/") {
		global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
		response.FailWithMessage(errors.New("error paths in request").Error(), c)
		return
	}
	dir := path.Dir(paths[0])
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
			response.FailWithMessage(errors.New("mkdir %s failed, err: %v").Error(), c)
			return
		}
	}
	success := 0
	failures := make(buserr.MultiErr)
	for _, file := range files {
		if err := c.SaveUploadedFile(file, path.Join(paths[0], file.Filename)); err != nil {
			e := fmt.Errorf("upload [%s] file failed, err: %v", file.Filename, err)
			failures[file.Filename] = e
			global.GVA_LOG.Error(e.Error())
			continue
		}
		success++
	}
	if success == 0 {
		response.FailWithMessage(failures.Error(), c)
	} else {
		response.OkWithMessage(fmt.Sprintf("%d files upload success", success), c)
	}
}

// @Tags File
// @Summary Check file exist
// @Description 检测文件是否存在
// @Accept json
// @Param request body request.FilePathCheck true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/check [post]
func (Frw *FileRWApi) CheckFile(c *gin.Context) {
	var req request.FilePathCheck
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if _, err := os.Stat(req.Path); err != nil {
		response.OkWithData(false, c)
		return
	}
	response.OkWithData(true, c)
}
