package example

import "github.com/Allen9012/ServerManager/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
