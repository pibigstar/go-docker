package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
/**
	network namespace 主要是用来隔离网络设备，IP地址端口等
	它可以让每个容器拥有自己独立的（虚拟的）网络设备
	每个namespace中的端口都不会互相冲突
 */

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err !=nil {
		log.Fatal(err)
	}
}
