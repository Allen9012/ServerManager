package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//var DefaultDockerClient

func ConnectDocker() {
	// 创建 Docker 客户端
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// 定义容器配置
	config := &container.Config{
		Image: "alpine:latest",                                // 使用 alpine 镜像作为示例
		Cmd:   []string{"/bin/sh", "-c", "ls -la /host/path"}, // 容器内执行的命令
	}

	// 定义主机配置，包括挂载的目录
	hostConfig := &container.HostConfig{
		Mounts: []types.Mount{
			{
				Type:   types.MountTypeBind, // 使用 bind 类型挂载
				Source: "/path/on/host",     // 宿主机上的路径
				Target: "/host/path",        // 容器内的路径
			},
		},
	}

	// 创建容器
	resp, err := cli.ContainerCreate(context.Background(), config, hostConfig, nil, "")
	if err != nil {
		panic(err)
	}

	// 启动容器
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// 等待容器停止
	w, err := cli.ContainerWait(context.Background(), resp.ID, container.WaitConditionNotRunning)
	if err != nil {
		panic(err)
	}

	// 输出容器日志
	if w.StatusCode != 0 {
		logs, err := cli.ContainerLogs(context.Background(), resp.ID, types.ContainerLogsOptions{ShowStdout: true})
		if err != nil {
			panic(err)
		}
		os.Stdout.Write(logs)
	}

	fmt.Printf("Container exited with status: %d\n", w.StatusCode)
}
