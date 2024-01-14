package service

import (
	"github.com/Allen9012/ServerManager/server/service/example"
	"github.com/Allen9012/ServerManager/server/service/file"
	"github.com/Allen9012/ServerManager/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	FileServiceGroup    file.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
