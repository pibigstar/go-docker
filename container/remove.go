package container

import (
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"go-docker/common"
)

func RemoveContainer(containerName string) {
	info, err := getContainerInfo(containerName)
	if err != nil {
		logrus.Errorf("get container info, err: %v", err)
		return
	}
	// 只能删除停止状态的容器
	if info.Status != common.Stop {
		logrus.Errorf("can't remove running container")
		return
	}
	dir := path.Join(common.DefaultContainerInfoPath, containerName)
	err = os.RemoveAll(dir)
	if err != nil {
		logrus.Errorf("remove container dir: %s, err: %v", dir, err)
		return
	}
}
