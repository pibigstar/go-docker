package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
user namespace 主要是用来隔离用户和用户组id的
https://github.com/xianlubird/mydocker/issues/3
echo 640 > /proc/sys/user/max_user_namespaces
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				// 容器的UID
				ContainerID: 1,
				// 宿主机的UID
				HostID: 0,
				Size:   1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				// 容器的GID
				ContainerID: 1,
				// 宿主机的GID
				HostID: 0,
				Size:   1,
			},
		},
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
