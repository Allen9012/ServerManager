package response

import "github.com/Allen9012/ServerManager/server/model/file"

type FileInfo struct {
	file.FileInfo
}

type UploadInfo struct {
	Name      string `json:"name"`
	Size      int    `json:"size"`
	CreatedAt string `json:"createdAt"`
}

type FileTree struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	Children []FileTree `json:"children"`
}

type DirSizeRes struct {
	Size float64 `json:"size" validate:"required"`
}

type FileProcessKeys struct {
	Keys []string `json:"keys"`
}

type FileWgetRes struct {
	Key string `json:"key"`
}

type FileLineContent struct {
	Content string `json:"content"`
	End     bool   `json:"end"`
	Path    string `json:"path"`
}

type FileExist struct {
	Exist bool `json:"exist"`
}
