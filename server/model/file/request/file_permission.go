package request

import (
	"github.com/Allen9012/ServerManager/server/model/common/request"
	"time"
	
)

type FilePermissionSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
