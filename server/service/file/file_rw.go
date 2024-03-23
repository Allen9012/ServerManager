package file

import (
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/file/request"
	"github.com/Allen9012/ServerManager/server/model/file/response"
	"github.com/Allen9012/ServerManager/server/utils/buserr"
	"github.com/Allen9012/ServerManager/server/utils/constant"
	"github.com/Allen9012/ServerManager/server/utils/files"
	"io/fs"
	"os"
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
