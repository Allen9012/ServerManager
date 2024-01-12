package file

import (
	"github.com/Allen9012/ServerManager/server/global"
)

type FilePermission struct {
	global.GVA_MODEL
	PermissionID    uint   `json:"permission_id" gorm:"not null;primary_key;comment:权限id"`
	PermissionState uint   `json:"permission_state" gorm:"not null;comment:权限状态 rw"`
	Regexp          string `json:"regexp" gorm:"regexp;comment:正则"`
}
