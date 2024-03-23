package files

import "github.com/spf13/afero"

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
