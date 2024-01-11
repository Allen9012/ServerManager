package response

import (
	"github.com/Allen9012/ServerManager/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
