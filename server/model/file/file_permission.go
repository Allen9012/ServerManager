// 自动生成模板FilePermission
package file

import (
	"github.com/Allen9012/ServerManager/server/global"
	
	
)

// 描述文件权限的信息 结构体  FilePermission
type FilePermission struct {
      global.GVA_MODEL
      PermissionState  *int `json:"permissionState" form:"permissionState" gorm:"column:permission_state;comment:;"binding:"required"`  //权限状态 rw 
      Regexp  string `json:"regexp" form:"regexp" gorm:"column:regexp;comment:;"`  //正则 
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:;"binding:"required"`  //用户id关联 
}


// TableName 描述文件权限的信息 FilePermission自定义表名 file_permission
func (FilePermission) TableName() string {
  return "file_permission"
}

