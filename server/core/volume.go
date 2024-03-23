package core

import (
	"errors"
	"fmt"
	"github.com/Allen9012/ServerManager/server/global"
	"go.uber.org/zap"
	"os"
	"strings"
)

type DockerVolume struct {
	HostPath      []string
	ContainerPath []string
}

// 添加path后主动点击更新
func (d *DockerVolume) UpdatePath() error {
	volume := global.GVA_CONFIG.Volume
	hpath := strings.Split(volume.HostPath, ";")
	cpath := strings.Split(volume.ContainerPath, ";")
	if len(hpath) != len(cpath) {
		return errors.New(fmt.Sprintf("error config volume :%+v", volume))
	}
	zap.L().Debug("update volume path", zap.Any("volume", volume))
	return nil
}

// 检查是否挂载目录
func checkMount() error {
	var err error
	volume := DockerVolume{}
	err = volume.UpdatePath()
	if err != nil {
		return err
	}
	// 查看是否挂载正确
	for _, path := range volume.ContainerPath {
		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			zap.L().Warn(fmt.Sprintf("run path in docker: %v", path))
			// 暂时不处理
		} else if err != nil {
			// 处理其他可能的错误
			zap.L().Warn(fmt.Sprintf("检查路径时发生错误:", err))
			return err
		} else {
			if info.IsDir() {
				zap.L().Info(fmt.Sprintf("run path in docker: %v", path))
			} else {
				zap.L().Info(fmt.Sprintf("not path : %v", path))
			}
		}
	}
	return nil
}
