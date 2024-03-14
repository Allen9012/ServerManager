package config

import (
	"github.com/Allen9012/ServerManager/server/global"
	"strings"
)

type Volume struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}

type DockerVolume struct {
	Path []string
}

// UpdatePath 添加path后主动点击更新
func (d *DockerVolume) UpdatePath() {
	volume := global.GVA_CONFIG.Volume
	paths := volume.Path
	path := strings.Split(paths, ";")
	d.Path = path
}
