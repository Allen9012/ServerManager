package request

import (
	"github.com/Allen9012/ServerManager/server/model/common/request"
	"github.com/Allen9012/ServerManager/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
