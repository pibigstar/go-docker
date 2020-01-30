package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"go-docker/cgroups"
	"go-docker/cgroups/subsystem"
	"go-docker/container"
)

func Run(cmdArray []string, tty bool, res *subsystem.ResourceConfig) error {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logrus.Errorf("failed to new parent process")
		return fmt.Errorf("failed to new parent process")
	}
	if err := parent.Start(); err !=nil {
		logrus.Errorf("parent start failed, err: %v", err)
		return err
	}
	// 添加资源限制
	cgroupMananger := cgroups.NewCGroupManager("go-docker-cgroup")
	// 删除资源限制
	defer cgroupMananger.Destroy()
	// 设置资源限制
	cgroupMananger.Set(res)
	// 将容器进程，加入到各个subsystem挂载对应的cgroup中
	cgroupMananger.Apply(parent.Process.Pid)

	sendInitCommand(cmdArray, writePipe)
	return parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}