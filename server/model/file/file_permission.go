// 自动生成模板FilePermission
package file

import (
	"github.com/Allen9012/ServerManager/server/global"
)

/*	存储使用	*/

// 描述文件权限的信息 结构体  FilePermission
type FilePermission struct {
	global.GVA_MODEL
	PermissionState int    `json:"permissionState" form:"permissionState" gorm:"column:permission_state;comment:;"` //权限状态 rw
	Regexp          string `json:"regexp" form:"regexp" gorm:"column:regexp;comment:;"`                             //正则
	UserId          int    `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`                            //角色id关联
}

// TableName 描述文件权限的信息 FilePermission自定义表名 file_permission
func (FilePermission) TableName() string {
	return "file_permission"
}

const (
	DENY int = iota
	R
	W
	RW
)

type RWAction int

const (
	Upload RWAction = iota
	Download
)
