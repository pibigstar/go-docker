package container

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/sirupsen/logrus"

	"go-docker/common"
)

// 创建一个会隔离namespace进程的Command
func NewParentProcess(tty bool, volume string) (*exec.Cmd, *os.File) {
	readPipe, writePipe, _ := os.Pipe()
	// 调用自身，传入 init 参数，也就是执行 initCommand
	cmd := exec.Command("/proc/self/exe", "init")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.ExtraFiles = []*os.File{
		readPipe,
	}
	err := NewWorkSpace(common.RootPath, common.MntPath, volume)
	if err != nil {
		logrus.Errorf("new work space, err: %v", err)
	}
	// 指定容器初始化后的工作目录
	cmd.Dir = common.MntPath
	return cmd, writePipe
}
