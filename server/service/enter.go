package service

import (
	"github.com/Allen9012/ServerManager/server/service/example"
	"github.com/Allen9012/ServerManager/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
