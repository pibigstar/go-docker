package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
/**
	mount namespace 用来隔离各个进程看到的挂载点视图
 */

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err !=nil {
		log.Fatal(err)
	}
}
