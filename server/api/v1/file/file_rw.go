package file

import (
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/common/response"
	"github.com/Allen9012/ServerManager/server/model/file/request"
	"github.com/Allen9012/ServerManager/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

type FileRWApi struct {
}

var FRWService = service.ServiceGroupApp.FileServiceGroup.FileRWService

// ListFiles
// @Tags File
// @Summary List files
// @Description 获取文件列表
// @Accept json
// @Param request body request.FileOption true "request"
// @Success 200 {object} response.FileInfo
// @Security ApiKeyAuth
// @Router /files/search [post]
func (frw *FileRWApi) ListFiles(c *gin.Context) {
	var req request.FileOption
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
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

// CreateFile
// @Tags File
// @Summary Create file
// @Description 创建文件/文件夹
// @Accept json
// @Param request body request.FileCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"创建文件/文件夹 [path]","formatEN":"Create dir or file [path]"}
func (frw *FileRWApi) CreateFile(c *gin.Context) {
	var req request.FileCreate
	var err error
	err = c.ShouldBindJSON(&req)
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

// DeleteFile
// @Tags File
// @Summary Delete file
// @Description 删除文件/文件夹
// @Accept json
// @Param request body request.FileDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/del [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"删除文件/文件夹 [path]","formatEN":"Delete dir or file [path]"}
func (frw *FileRWApi) DeleteFile(c *gin.Context) {
	var req request.FileDelete
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FRWService.Delete(req)
	if err != nil {
		global.GVA_LOG.Error("DeleteFile fail", zap.Error(err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// @Tags File
// @Summary Batch delete file
// @Description 批量删除文件/文件夹
// @Accept json
// @Param request body request.FileBatchDelete true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/batch/del [post]
// @x-panel-log {"bodyKeys":["paths"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"批量删除文件/文件夹 [paths]","formatEN":"Batch delete dir or file [paths]"}
func (frw *FileRWApi) BatchDeleteFile(c *gin.Context) {
	var req request.FileBatchDelete
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FRWService.BatchDelete(req)
	if err != nil {
		global.GVA_LOG.Error("BatchDeleteFile fail", zap.Error(err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// UploadFiles
// @Tags File
// @Summary Upload file
// @Description 上传文件
// @Param file formData file true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/upload [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"上传文件 [path]","formatEN":"Upload file [path]"}
// TODO modify logic
//func (Frw *FileRWApi) UploadFiles(c *gin.Context) {
//	form, err := c.MultipartForm()
//	if err != nil {
//		global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	files := form.File["file"]
//	paths := form.Value["path"]
//	if len(paths) == 0 || !strings.Contains(paths[0], "/") {
//		global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
//		response.FailWithMessage(errors.New("error paths in request").Error(), c)
//		return
//	}
//	dir := path.Dir(paths[0])
//	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
//		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
//			global.GVA_LOG.Error("UploadFiles fail", zap.Error(err))
//			response.FailWithMessage(errors.New("mkdir %s failed, err: %v").Error(), c)
//			return
//		}
//	}
//	success := 0
//	failures := make(buserr.MultiErr)
//	for _, file := range files {
//		if err := c.SaveUploadedFile(file, path.Join(paths[0], file.Filename)); err != nil {
//			e := fmt.Errorf("upload [%s] file failed, err: %v", file.Filename, err)
//			failures[file.Filename] = e
//			global.GVA_LOG.Error(e.Error())
//			continue
//		}
//		success++
//	}
//	if success == 0 {
//		response.FailWithMessage(failures.Error(), c)
//	} else {
//		response.OkWithMessage(fmt.Sprintf("%d files upload success", success), c)
//	}
//}

// CheckFile
// @Tags File
// @Summary Check file exist
// @Description 检测文件是否存在
// @Accept json
// @Param request body request.FilePathCheck true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/check [post]
func (frw *FileRWApi) CheckFile(c *gin.Context) {
	var req request.FilePathCheck
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if _, err = os.Stat(req.Path); err != nil {
		response.OkWithData(false, c)
		return
	}
	response.OkWithData(true, c)
}

// @Tags File
// @Summary Move file
// @Description 移动文件
// @Accept json
// @Param request body request.FileMove true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/move [post]
// @x-panel-log {"bodyKeys":["oldPaths","newPath"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"移动文件 [oldPaths] => [newPath]","formatEN":"Move [oldPaths] => [newPath]"}
func (frw *FileRWApi) MoveFile(c *gin.Context) {
	var req request.FileMove
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = FRWService.MvFile(req); err != nil {
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(true, c)
}

// @Tags File
// @Summary Change file name
// @Description 修改文件名称
// @Accept json
// @Param request body request.FileRename true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/rename [post]
// @x-panel-log {"bodyKeys":["oldName","newName"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"重命名 [oldName] => [newName]","formatEN":"Rename [oldName] => [newName]"}
func (frw *FileRWApi) ChangeFileName(c *gin.Context) {
	var req request.FileRename
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = FRWService.ChangeName(req); err != nil {
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(true, c)
}
