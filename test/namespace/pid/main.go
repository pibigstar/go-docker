package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
pid namespace 主要是用来隔离 进程ID的
同样一个进程在不同的pid namespace中可以拥有不同的PID
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
