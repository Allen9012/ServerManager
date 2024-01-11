package response

import "github.com/Allen9012/ServerManager/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
