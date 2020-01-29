package container

import (
	"github.com/sirupsen/logrus"
	"os"
	"syscall"
)

// 本容器执行的第一个进程
// 使用mount挂载proc文件系统
// 以便后面通过`ps`等系统命令查看当前进程资源的情况
func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)
	// 挂载
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		return err
	}

	err = syscall.Exec(command, []string{command}, os.Environ())
	if err != nil {
		return err
	}
	return nil
}