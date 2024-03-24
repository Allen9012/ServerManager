package files

import (
	"fmt"
	"github.com/Allen9012/ServerManager/server/utils/cmd"
	"github.com/spf13/afero"
	"io/fs"
	"os"
	"path"
	"path/filepath"
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

// 剪切文件
func (f FileOp) Cut(oldPaths []string, dst, name string, cover bool) error {
	for _, p := range oldPaths {
		var dstPath string
		if name != "" {
			dstPath = filepath.Join(dst, name)
			if f.Stat(dstPath) {
				dstPath = dst
			}
		} else {
			// 不存在的路径就在基路径创建一个
			base := filepath.Base(p)
			dstPath = filepath.Join(dst, base)
		}
		coverFlag := ""
		if cover {
			coverFlag = "-f"
		}

		cmdStr := fmt.Sprintf(`mv %s "%s" "%s"`, coverFlag, p, dstPath)
		if err := cmd.ExecCmd(cmdStr); err != nil {
			return err
		}
	}
	return nil
}

func (f FileOp) CopyAndReName(src, dst, name string, cover bool) error {
	// 路径有问题
	if src = path.Clean("/" + src); src == "" {
		return os.ErrNotExist
	}
	if dst = path.Clean("/" + dst); dst == "" {
		return os.ErrNotExist
	}
	if src == "/" || dst == "/" {
		return os.ErrInvalid
	}
	if dst == src {
		return os.ErrInvalid
	}

	srcInfo, err := f.Fs.Stat(src)
	if err != nil {
		return err
	}

	if srcInfo.IsDir() {
		dstPath := dst
		if name != "" && !cover {
			dstPath = filepath.Join(dst, name)
		}
		return cmd.ExecCmd(fmt.Sprintf(`cp -rf "%s" "%s"`, src, dstPath))
	} else {
		dstPath := filepath.Join(dst, name)
		if cover {
			dstPath = dst
		}
		return cmd.ExecCmd(fmt.Sprintf(`cp -f "%s" "%s"`, src, dstPath))
	}
}

func (f FileOp) Rename(oldName string, newName string) error {
	return f.Fs.Rename(oldName, newName)
}
