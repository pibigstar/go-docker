package container

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/sirupsen/logrus"

	"go-docker/common"
)

// 创建容器运行时目录
func NewWorkSpace(volume, containerName, imageName string) error {
	// 1. 创建只读层
	err := createReadOnlyLayer(imageName)
	if err != nil {
		logrus.Errorf("create read only layer, err: %v", err)
		return err
	}
	// 2. 创建读写层
	err = createWriteLayer(containerName)
	if err != nil {
		logrus.Errorf("create write layer, err: %v", err)
		return err
	}
	// 3. 创建挂载点，将只读层和读写层挂载到指定位置
	err = CreateMountPoint(containerName, imageName)
	if err != nil {
		logrus.Errorf("create mount point, err: %v", err)
		return err
	}
	// 4. 设置宿主机与容器文件映射
	mountVolume(containerName, imageName, volume)
	return nil
}

// 根据镜像创建只读层
func createReadOnlyLayer(imageName string) error {
	imagePath := path.Join(common.RootPath, imageName)
	_, err := os.Stat(imagePath)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(imagePath, os.ModePerm)
		if err != nil {
			logrus.Errorf("mkdir image path, err: %v", err)
			return err
		}
	}
	// 解压 /root/imageName.tar
	imageTarPath := path.Join(common.RootPath, fmt.Sprintf("%s.tar", imageName))
	if _, err = exec.Command("tar", "-xvf", imageTarPath, "-C", imagePath).CombinedOutput(); err != nil {
		logrus.Errorf("tar image tar,path: %s, err: %v", imageTarPath, err)
		return err
	}
	return nil
}

// 创建读写层
func createWriteLayer(containerName string) error {
	writeLayerPath := path.Join(common.RootPath, common.WriteLayer, containerName)
	_, err := os.Stat(writeLayerPath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(writeLayerPath, os.ModePerm)
		if err != nil {
			logrus.Errorf("mkdir write layer, err: %v", err)
			return err
		}
	}
	return nil
}

func CreateMountPoint(containerName, imageName string) error {
	mntPath := path.Join(common.MntPath, containerName)
	_, err := os.Stat(mntPath)
	if err != nil && os.IsNotExist(err) {
		err := os.MkdirAll(mntPath, os.ModePerm)
		if err != nil {
			logrus.Errorf("mkdir mnt path, err: %v", err)
			return err
		}
	}

	// 将宿主机上关于容器的读写层和只读层挂载到 /root/mnt/容器名 里
	writeLayPath := path.Join(common.RootPath, common.WriteLayer, containerName)
	imagePath := path.Join(common.RootPath, imageName)
	dirs := fmt.Sprintf("dirs=%s:%s", writeLayPath, imagePath)
	cmd := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", mntPath)
	if err := cmd.Run(); err != nil {
		logrus.Errorf("mnt cmd run, err: %v", err)
		return err
	}
	return nil
}

func mountVolume(containerName, imageName, volume string) {
	if volume != "" {
		volumes := strings.Split(volume, ":")
		if len(volumes) > 1 {
			// 创建宿主机中文件路径
			parentPath := volumes[0]
			if _, err := os.Stat(parentPath); err != nil && os.IsNotExist(err) {
				if err := os.MkdirAll(parentPath, os.ModePerm); err != nil {
					logrus.Errorf("mkdir parent path: %s, err: %v", parentPath, err)
				}
			}

			// 创建容器内挂载点
			containerPath := volumes[1]
			containerVolumePath := path.Join(common.MntPath, containerName, containerPath)
			if _, err := os.Stat(containerVolumePath); err != nil && os.IsNotExist(err) {
				if err := os.MkdirAll(containerVolumePath, os.ModePerm); err != nil {
					logrus.Errorf("mkdir volume path path: %s, err: %v", containerVolumePath, err)
				}
			}

			// 把宿主机文件目录挂载到容器挂载点中
			dirs := fmt.Sprintf("dirs=%s", parentPath)
			cmd := exec.Command("mount", "-t", "aufs", "-o", dirs, "none", containerVolumePath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				logrus.Errorf("mount cmd run, err: %v", err)
			}
		}
	}
}

// 删除容器工作空间
func DeleteWorkSpace(containerName, volume string) error {
	// 1. 卸载挂载点
	err := unMountPoint(containerName)
	if err != nil {
		return err
	}
	// 2. 删除读写层
	err = deleteWriteLayer(containerName)
	if err != nil {
		return err
	}
	// 3. 删除宿主机与文件系统映射
	deleteVolume(containerName, volume)
	return nil
}

func unMountPoint(containerName string) error {
	mntPath := path.Join(common.MntPath, containerName)
	if _, err := exec.Command("umount", mntPath).CombinedOutput(); err != nil {
		logrus.Errorf("unmount mnt, err: %v", err)
		return err
	}
	err := os.RemoveAll(mntPath)
	if err != nil {
		logrus.Errorf("remove mnt path, err: %v", err)
		return err
	}
	return nil
}

func deleteWriteLayer(containerName string) error {
	writerLayerPath := path.Join(common.RootPath, common.WriteLayer, containerName)
	return os.RemoveAll(writerLayerPath)
}

func deleteVolume(containerName, volume string) {
	if volume != "" {
		volumes := strings.Split(volume, ":")
		if len(volumes) > 1 {
			mntPath := path.Join(common.MntPath, common.WriteLayer, containerName)
			containerPath := path.Join(mntPath, volumes[1])
			if _, err := exec.Command("umount", containerPath).CombinedOutput(); err != nil {
				logrus.Errorf("unmount container path, err: %v", err)
			}
		}
	}
}
