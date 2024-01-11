package service

 {{- if .NeedModel }}
import (
   "github.com/Allen9012/ServerManager/server/plugin/{{ .Snake}}/model"
)
{{ end }}

type {{ .PlugName}}Service struct{}

func (e *{{ .PlugName}}Service) PlugService({{- if .HasRequest }}req model.Request {{ end -}}) ({{- if .HasResponse }}res model.Response,{{ end -}} err error) {
    // 写你的业务逻辑
	return {{- if .HasResponse }} res,{{ end }} nil
}
