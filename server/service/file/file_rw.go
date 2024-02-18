package file

import (
	"github.com/Allen9012/ServerManager/server/model/common/request"
	"github.com/Allen9012/ServerManager/server/model/common/response"
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
	//fo := file.NewFileOp()
	//if fo.Stat(op.Path) {
	//	return buserr.New(constant.ErrFileIsExit)
	//}
	//if op.IsDir {
	//	return fo.CreateDir(op.Path, fs.FileMode(op.Mode))
	//} else {
	//	if op.IsLink {
	//		if !fo.Stat(op.LinkPath) {
	//			return buserr.New(constant.ErrLinkPathNotFound)
	//		}
	//		return fo.LinkFile(op.LinkPath, op.Path, op.IsSymlink)
	//	} else {
	//		return fo.CreateFile(op.Path)
	//	}
	//}
	panic("implement me")
}

func (f *FileRWService) GetFileList(op request.FileOption) (response.FileInfo, error) {
	//var fileInfo response.FileInfo
	//if _, err := os.Stat(op.Path); err != nil && os.IsNotExist(err) {
	//	return fileInfo, nil
	//}
	//info, err := file.NewFileInfo(op.FileOption)
	//if err != nil {
	//	return fileInfo, err
	//}
	//fileInfo.FileInfo = *info
	//return fileInfo, nil
	panic("implement me")
}
