package file

import (
	"github.com/spf13/afero"
	"io/fs"
	"os"
	"time"
)

type FileInfo struct {
	Fs         afero.Fs    `json:"-"`
	Path       string      `json:"path"`
	Name       string      `json:"name"`
	User       string      `json:"user"`
	Group      string      `json:"group"`
	Uid        string      `json:"uid"`
	Gid        string      `json:"gid"`
	Extension  string      `json:"extension"`
	Content    string      `json:"content"`
	Size       int64       `json:"size"`
	IsDir      bool        `json:"isDir"`
	IsSymlink  bool        `json:"isSymlink"`
	IsHidden   bool        `json:"isHidden"`
	LinkPath   string      `json:"linkPath"`
	Type       string      `json:"type"`
	Mode       string      `json:"mode"`
	MimeType   string      `json:"mimeType"`
	UpdateTime time.Time   `json:"updateTime"`
	ModTime    time.Time   `json:"modTime"`
	FileMode   os.FileMode `json:"-"`
	Items      []*FileInfo `json:"items"`
	ItemTotal  int         `json:"itemTotal"`
	FavoriteID uint        `json:"favoriteID"`
}

type FileOption struct {
	Path       string `json:"path"`
	Search     string `json:"search"`
	ContainSub bool   `json:"containSub"`
	Expand     bool   `json:"expand"`
	Dir        bool   `json:"dir"`
	ShowHidden bool   `json:"showHidden"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
	SortBy     string `json:"sortBy"`
	SortOrder  string `json:"sortOrder"`
}

type FileSearchInfo struct {
	Path string `json:"path"`
	fs.FileInfo
}
