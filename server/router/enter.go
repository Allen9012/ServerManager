package router

import (
	"github.com/Allen9012/ServerManager/server/router/example"
	"github.com/Allen9012/ServerManager/server/router/file"
	"github.com/Allen9012/ServerManager/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	File    file.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
