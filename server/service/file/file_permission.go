package file

import (
	"github.com/Allen9012/ServerManager/server/global"
	"github.com/Allen9012/ServerManager/server/model/file"
	"github.com/Allen9012/ServerManager/server/model/common/request"
    fileReq "github.com/Allen9012/ServerManager/server/model/file/request"
)

type FilePermissionService struct {
}

// CreateFilePermission 创建描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService) CreateFilePermission(FP *file.FilePermission) (err error) {
	err = global.GVA_DB.Create(FP).Error
	return err
}

// DeleteFilePermission 删除描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService)DeleteFilePermission(FP file.FilePermission) (err error) {
	err = global.GVA_DB.Delete(&FP).Error
	return err
}

// DeleteFilePermissionByIds 批量删除描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService)DeleteFilePermissionByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]file.FilePermission{},"id in ?",ids.Ids).Error
	return err
}

// UpdateFilePermission 更新描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService)UpdateFilePermission(FP file.FilePermission) (err error) {
	err = global.GVA_DB.Save(&FP).Error
	return err
}

// GetFilePermission 根据id获取描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService)GetFilePermission(id uint) (FP file.FilePermission, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FP).Error
	return
}

// GetFilePermissionInfoList 分页获取描述文件权限的信息记录
// Author Allen
func (FPService *FilePermissionService)GetFilePermissionInfoList(info fileReq.FilePermissionSearch) (list []file.FilePermission, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&file.FilePermission{})
    var FPs []file.FilePermission
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&FPs).Error
	return  FPs, total, err
}
