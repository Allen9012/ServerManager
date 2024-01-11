package global

{{- if .HasGlobal }}

import "github.com/Allen9012/ServerManager/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}