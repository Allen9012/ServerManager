package files

import (
	"github.com/spf13/afero"
	"io/fs"
	"os"
)

type FileOp struct {
	Fs afero.Fs
}

func NewFileOp() FileOp {
	return FileOp{
		Fs: afero.NewOsFs(),
	}
}

// 获取文件os.stat 封装
func (f FileOp) Stat(dst string) bool {
	info, _ := f.Fs.Stat(dst)
	return info != nil
}

func (f FileOp) CreateDir(dst string, mode fs.FileMode) error {
	return f.Fs.MkdirAll(dst, mode)
}

func (f FileOp) LinkFile(source string, dst string, isSymlink bool) error {
	if isSymlink {
		osFs := afero.OsFs{}
		return osFs.SymlinkIfPossible(source, dst)
	} else {
		return os.Link(source, dst)
	}
}

// 创建文件
func (f FileOp) CreateFile(dst string) error {
	if _, err := f.Fs.Create(dst); err != nil {
		return err
	}
	return nil
}

// 删除文件夹
func (f FileOp) DeleteDir(dst string) error {
	return f.Fs.RemoveAll(dst)
}

// 删除文件
func (f FileOp) DeleteFile(dst string) error {
	return f.Fs.Remove(dst)
}
