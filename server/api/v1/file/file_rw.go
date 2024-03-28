package file

import (
	"fmt"
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/common/response"
	"github.com/Allen9012/ServerManager/server/model/file"
	"github.com/Allen9012/ServerManager/server/model/file/request"
	"github.com/Allen9012/ServerManager/server/service"
	"github.com/Allen9012/ServerManager/server/utils"
	"github.com/Allen9012/ServerManager/server/utils/buserr"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
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
		global.GVA_LOG.Error(fmt.Sprintf("MoveFile fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
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
		global.GVA_LOG.Error(fmt.Sprintf("ChangeFileName fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(true, c)
}

// @Tags File
// @Summary Change file owner
// @Description 修改文件用户/组
// @Accept json
// @Param request body request.FileRoleUpdate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/owner [post]
// @x-panel-log {"bodyKeys":["path","user","group"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"修改用户/组 [paths] => [user]/[group]","formatEN":"Change owner [paths] => [user]/[group]"}
func (frw *FileRWApi) ChangeFileOwner(c *gin.Context) {
	var req request.FileRoleUpdate
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := FRWService.ChangeOwner(req); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("ChangeFileOwner fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// @Tags File
// @Summary Change file mode
// @Description 修改文件权限
// @Accept json
// @Param request body request.FileCreate true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/mode [post]
// @x-panel-log {"bodyKeys":["path","mode"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"修改权限 [paths] => [mode]","formatEN":"Change mode [paths] => [mode]"}
func (frw *FileRWApi) ChangeFileMode(c *gin.Context) {
	var req request.FileCreate
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = FRWService.ChangeMode(req)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("ChangeFileMode fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// @Tags File
// @Summary Batch change file mode and owner
// @Description 批量修改文件权限和用户/组
// @Accept json
// @Param request body request.FileRoleReq true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/batch/role [post]
// @x-panel-log {"bodyKeys":["paths","mode","user","group"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"批量修改文件权限和用户/组 [paths] => [mode]/[user]/[group]","formatEN":"Batch change file mode and owner [paths] => [mode]/[user]/[group]"}
func (frw *FileRWApi) BatchChangeModeAndOwner(c *gin.Context) {
	var req request.FileRoleReq
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := FRWService.BatchChangeModeAndOwner(req); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("BatchChangeModeAndOwner fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(nil, c)
}

// @Tags File
// @Summary Load file content
// @Description 获取文件内容
// @Accept json
// @Param request body request.FileContentReq true "request"
// @Success 200 {object} response.FileInfo
// @Security ApiKeyAuth
// @Router /files/content [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"获取文件内容 [path]","formatEN":"Load file content [path]"}
func (frw *FileRWApi) GetContent(c *gin.Context) {
	var req request.FileContentReq
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := FRWService.GetContent(req)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("GetContent fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(info, c)
}

// @Tags File
// @Summary Page file
// @Description 分页获取上传文件
// @Accept json
// @Param request body request.SearchUploadWithPage true "request"
// @Success 200 {array} response.FileInfo
// @Security ApiKeyAuth
// @Router /files/upload/search [post]
func (frw *FileRWApi) SearchUploadWithPage(c *gin.Context) {
	var req request.SearchUploadWithPage
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	total, files, err := FRWService.SearchUploadWithPage(req)
	if err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("SearchUploadWithPage fail: %v", err))
		response.FailWithMessage("ErrInternalServer", c)
		return
	}
	response.OkWithData(request.PageResult{
		Items: files,
		Total: total,
	}, c)
}

// @Tags File
// @Summary Download file
// @Description 下载文件
// @Accept json
// @Param request body request.FileDownload true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/download [get]
// @x-panel-log {"bodyKeys":["name"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"下载文件 [name]","formatEN":"Download file [name]"}
func (frw *FileRWApi) Download(c *gin.Context) {
	filePath := c.Query("path")
	f, err := os.Open(filePath)
	if err != nil {
		response.FailWithMessage("ErrTypeInvalidParams", c)
		return
	}
	id := utils.GetUserAuthorityId(c)
	global.GVA_LOG.Debug(fmt.Sprintf("Download filePath: %v", filePath))
	if ok := FRWService.CheckPermission(id, filePath, file.Download); !ok {
		global.GVA_LOG.Error(fmt.Sprintf("Download Permission Denied authid: %d, filePath: %s", id, filePath))
		response.FailWithMessage("Err Permission Denied", c)
		return
	}
	info, _ := f.Stat()
	c.Header("Content-Length", strconv.FormatInt(info.Size(), 10))
	c.Header("Content-Disposition", "attachment; filename*=utf-8''"+url.PathEscape(info.Name()))
	http.ServeContent(c.Writer, c.Request, info.Name(), info.ModTime(), f)
}

// @Tags File
// @Summary Upload file
// @Description 上传文件
// @Param file formData file true "request"
// @Success 200
// @Security ApiKeyAuth
// @Router /files/upload [post]
// @x-panel-log {"bodyKeys":["path"],"paramKeys":[],"BeforeFunctions":[],"formatZH":"上传文件 [path]","formatEN":"Upload file [path]"}
func (frw *FileRWApi) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage("ErrTypeInvalidParams", c)
		return
	}
	// 权限校验
	id := utils.GetUserAuthorityId(c)

	files := form.File["file"]
	paths := form.Value["path"]
	if len(paths) == 0 || !strings.Contains(paths[0], "/") {
		response.FailWithDetailed(gin.H{}, "error paths in request", c)
		return
	}
	dir := path.Dir(paths[0])
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			response.FailWithDetailed(gin.H{}, fmt.Sprintf("mkdir %s failed, err: %v", dir, err), c)
			return
		}
	}
	// 上传的时候检查的是dir是否有没有权限，所以没有权限所有的文件都不能上传
	global.GVA_LOG.Debug(fmt.Sprintf("upload dir: %v", dir))
	if ok := FRWService.CheckPermission(id, dir, file.Upload); !ok {
		global.GVA_LOG.Error(fmt.Sprintf("Download Permission Denied authid: %d, dir: %s", id, dir))
		response.FailWithMessage("Err Permission Denied", c)
		return
	}
	// 记录上传失败的文件
	success := 0
	failures := make(buserr.MultiErr)
	for _, f := range files {
		if err := c.SaveUploadedFile(f, path.Join(paths[0], f.Filename)); err != nil {
			e := fmt.Errorf("upload [%s] file failed, err: %v", f.Filename, err)
			failures[f.Filename] = e
			global.GVA_LOG.Error(fmt.Sprintf("upload [%s] file failed, err: %v", f.Filename, err))
			continue
		}
		success++
	}
	if success == 0 {
		response.FailWithDetailed(failures, "ErrTypeInternalServer", c)
	} else {
		response.OkWithMessage(fmt.Sprintf("%d files upload success", success), c)
	}
}
