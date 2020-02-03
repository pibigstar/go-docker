package container

import (
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"

	"go-docker/common"
)

// 重新进入容器
// 通过设置环境变量的方式，让 C 语言写的程序真正执行
// 通过 setns 的系统调用，重新进入到指定的 PID的 namespace 中
func ExecContainer(containerName string, cmdArray []string) {
	info, err := getContainerInfo(containerName)
	if err != nil {
		logrus.Errorf("get container info, err: %v", err)
	}
	cmd := exec.Command("/proc/self/exe", "exec")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = os.Setenv(common.EnvExecPid, info.Pid)
	err = os.Setenv(common.EnvExecCmd, strings.Join(cmdArray, " "))
	if err = cmd.Run(); err != nil {
		logrus.Errorf("exec cmd run, err: %v", err)
	}
}
