package file

import (
	"errors"
	"fmt"
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/file"
	"github.com/Allen9012/ServerManager/server/model/file/request"
	"github.com/Allen9012/ServerManager/server/model/file/response"
	"github.com/Allen9012/ServerManager/server/utils/buserr"
	"github.com/Allen9012/ServerManager/server/utils/constant"
	"github.com/Allen9012/ServerManager/server/utils/files"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileRWService struct {
}

func (f *FileRWService) GetFileTree(op request.FileOption) ([]response.FileTree, error) {
	//var treeArray []response.FileTree
	//info, err := files.NewFileInfo(op.FileOption)
	//if err != nil {
	//	return nil, err
	//}
	//node := response.FileTree{
	//	ID:   common.GetUuid(),
	//	Name: info.Name,
	//	Path: info.Path,
	//}
	//for _, v := range info.Items {
	//	if v.IsDir {
	//		node.Children = append(node.Children, response.FileTree{
	//			ID:   common.GetUuid(),
	//			Name: v.Name,
	//			Path: v.Path,
	//		})
	//	}
	//}
	//return append(treeArray, node), nil
	panic("implement me")
}

func (f *FileRWService) Create(op request.FileCreate) error {
	fo := files.NewFileOp()
	if fo.Stat(op.Path) {
		return buserr.New(constant.ErrFileIsExit)
	}
	if op.IsDir {
		global.GVA_LOG.Info("2")
		return fo.CreateDir(op.Path, fs.FileMode(op.Mode))
	} else {
		// 创建文件
		if op.IsLink {
			if !fo.Stat(op.LinkPath) {
				return buserr.New(constant.ErrLinkPathNotFound)
			}
			return fo.LinkFile(op.LinkPath, op.Path, op.IsSymlink)
		} else {
			return fo.CreateFile(op.Path)
		}
	}
}

func (f *FileRWService) GetFileList(op request.FileOption) (response.FileInfo, error) {
	var fileInfo response.FileInfo
	// 没有文件返回空结构体
	if _, err := os.Stat(op.Path); err != nil && os.IsNotExist(err) {
		return fileInfo, nil
	}
	info, err := files.NewFileInfo(op.FileOption)
	if err != nil {
		return fileInfo, err
	}
	fileInfo.FileInfo = *info
	return fileInfo, nil
}

func (f *FileRWService) Delete(op request.FileDelete) error {
	fo := files.NewFileOp()
	//recycleBinStatus, _ := settingRepo.Get(settingRepo.WithByKey("FileRecycleBin"))
	//if recycleBinStatus.Value == "disable" {
	op.ForceDelete = true
	//}
	if op.ForceDelete {
		if op.IsDir {
			return fo.DeleteDir(op.Path)
		} else {
			return fo.DeleteFile(op.Path)
		}
	}
	//if err := NewIRecycleBinService().Create(request.RecycleBinCreate{SourcePath: op.Path}); err != nil {
	//	return err
	//}
	//return favoriteRepo.Delete(favoriteRepo.WithByPath(op.Path))
	return nil
}

// 批量删除
func (f *FileRWService) BatchDelete(op request.FileBatchDelete) error {
	fo := files.NewFileOp()
	if op.IsDir {
		for _, file := range op.Paths {
			if err := fo.DeleteDir(file); err != nil {
				return err
			}
		}
	} else {
		for _, file := range op.Paths {
			if err := fo.DeleteFile(file); err != nil {
				return err
			}
		}
	}
	return nil
}

func (f *FileRWService) MvFile(m request.FileMove) error {
	fo := files.NewFileOp()
	if !fo.Stat(m.NewPath) {
		return buserr.New(constant.ErrPathNotFound)
	}
	for _, oldPath := range m.OldPaths {
		// 找不到源文件
		if !fo.Stat(oldPath) {
			return buserr.WithName(constant.ErrFileNotFound, oldPath)
		}
		// 非法移动
		if oldPath == m.NewPath || strings.Contains(m.NewPath, filepath.Clean(oldPath)+"/") {
			return buserr.New(constant.ErrMovePathFailed)
		}
	}
	// 剪切
	if m.Type == "cut" {
		return fo.Cut(m.OldPaths, m.NewPath, m.Name, m.Cover)
	}
	// 拷贝
	var errs []error
	if m.Type == "copy" {
		for _, src := range m.OldPaths {
			if err := fo.CopyAndReName(src, m.NewPath, m.Name, m.Cover); err != nil {
				errs = append(errs, err)
				global.GVA_LOG.Error(fmt.Sprintf("copy file [%s] to [%s] failed, err: %s", src, m.NewPath, err.Error()))
			}
		}
	}

	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}
	if errString != "" {
		return errors.New(errString)
	}
	return nil
}

func (f *FileRWService) ChangeName(req request.FileRename) error {
	fo := files.NewFileOp()
	return fo.Rename(req.OldName, req.NewName)
}

func (f *FileRWService) ChangeOwner(req request.FileRoleUpdate) error {
	fo := files.NewFileOp()
	return fo.ChownR(req.Path, req.User, req.Group, req.Sub)
}

func (f *FileRWService) ChangeMode(op request.FileCreate) error {
	fo := files.NewFileOp()
	return fo.ChmodR(op.Path, op.Mode, op.Sub)
}

func (f *FileRWService) BatchChangeModeAndOwner(op request.FileRoleReq) error {
	fo := files.NewFileOp()
	for _, path := range op.Paths {
		if !fo.Stat(path) {
			return buserr.New(constant.ErrPathNotFound)
		}
		if err := fo.ChownR(path, op.User, op.Group, op.Sub); err != nil {
			return err
		}
		if err := fo.ChmodR(path, op.Mode, op.Sub); err != nil {
			return err
		}
	}
	return nil
}

func (f *FileRWService) GetContent(op request.FileContentReq) (response.FileInfo, error) {
	info, err := files.NewFileInfo(files.FileOption{
		Path:   op.Path,
		Expand: true,
	})
	if err != nil {
		return response.FileInfo{}, err
	}
	return response.FileInfo{FileInfo: *info}, nil
}

func (f *FileRWService) CheckPermission(authID uint, filePath string, flag file.RWAction) bool {
	if authID == constant.Admin {
		return true
	}
	// 取出所有和ID相关的regexp
	var permissions []file.FilePermission
	result := global.GVA_DB.Where("user_id = ?", authID).Find(&permissions)
	if result.RowsAffected == 0 {
		global.GVA_LOG.Error("no permission")
		return false
	}
	return _checkPermission(permissions, filePath, flag)
}

func _checkPermission(permissions []file.FilePermission, filePath string, flag file.RWAction) bool {
	for _, permission := range permissions {
		if flag == file.Download && (permission.PermissionState != file.R && permission.PermissionState != file.RW) {
			// 读
			continue
		} else if flag == file.Upload && (permission.PermissionState != file.W && permission.PermissionState != file.RW) {
			// 写
			continue
		} else {
			re := regexp.MustCompile(permission.Regexp)
			if re.MatchString(filePath) {
				//global.GVA_LOG.Debug(fmt.Sprintf("permission match, reg: %s, filePath: %s", permission.Regexp, filePath))
				return true
			} else {
				continue
			}
		}
	}
	// 对比当前文件是否满足权限
	return false
}

func (f *FileRWService) SearchUploadWithPage(req request.SearchUploadWithPage) (int64, interface{}, error) {
	var (
		files    []response.UploadInfo
		backData []response.UploadInfo
	)
	_ = filepath.Walk(req.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			files = append(files, response.UploadInfo{
				CreatedAt: info.ModTime().Format("2006-01-02 15:04:05"),
				Size:      int(info.Size()),
				Name:      info.Name(),
			})
		}
		return nil
	})
	total, start, end := len(files), (req.Page-1)*req.PageSize, req.Page*req.PageSize
	if start > total {
		backData = make([]response.UploadInfo, 0)
	} else {
		if end >= total {
			end = total
		}
		backData = files[start:end]
	}
	return int64(total), backData, nil
}
