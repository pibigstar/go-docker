package main

import (
	"go-docker/common"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"go-docker/cgroups"
	"go-docker/cgroups/subsystem"
	"go-docker/container"
)

func Run(cmdArray []string, tty bool, res *subsystem.ResourceConfig, volume, containerName string) {
	id := container.GenContainerID(10)
	if containerName == "" {
		containerName = id
	}
	parent, writePipe := container.NewParentProcess(tty, volume, containerName)
	if parent == nil {
		logrus.Errorf("failed to new parent process")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Errorf("parent start failed, err: %v", err)
		return
	}
	// 记录容器信息
	err := container.RecordContainerInfo(parent.Process.Pid, cmdArray, containerName, id)
	if err != nil {
		logrus.Errorf("record container info, err: %v", err)
	}

	// 添加资源限制
	cgroupMananger := cgroups.NewCGroupManager("go-docker")
	// 删除资源限制
	defer cgroupMananger.Destroy()
	// 设置资源限制
	cgroupMananger.Set(res)
	// 将容器进程，加入到各个subsystem挂载对应的cgroup中
	cgroupMananger.Apply(parent.Process.Pid)

	// 设置初始化命令
	sendInitCommand(cmdArray, writePipe)

	if tty {
		// 等待父进程结束
		err := parent.Wait()
		if err != nil {
			logrus.Errorf("parent wait, err: %v", err)
		}
		// 删除容器工作空间
		err = container.DeleteWorkSpace(common.RootPath, common.MntPath, volume)
		if err != nil {
			logrus.Errorf("delete work space, err: %v", err)
		}
		// 删除容器信息
		container.DeleteContainerInfo(containerName)
	}
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logrus.Infof("command all is %s", command)
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}
