package request

import (
	"github.com/Allen9012/ServerManager/server/model/common/request"
	"github.com/Allen9012/ServerManager/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
