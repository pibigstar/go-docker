package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
uts namespace 主要是用来隔离 nodename和 domainname 两个系统标识的
每个uts中允许有自己的 hostname
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
