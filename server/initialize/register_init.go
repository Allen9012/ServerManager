package initialize

import (
	_ "github.com/Allen9012/ServerManager/server/source/example"
	_ "github.com/Allen9012/ServerManager/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
