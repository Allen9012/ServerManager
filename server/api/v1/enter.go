package v1

import (
	"github.com/Allen9012/ServerManager/server/api/v1/example"
	"github.com/Allen9012/ServerManager/server/api/v1/file"
	"github.com/Allen9012/ServerManager/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	FileApiGroup    file.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
