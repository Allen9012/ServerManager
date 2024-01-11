package response

import "github.com/Allen9012/ServerManager/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
