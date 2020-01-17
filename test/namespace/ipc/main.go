package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
/**
	ipc namespace 主要是用来隔离 System VIPC 和 POSIX message queues的
 */
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err !=nil {
		log.Fatal(err)
	}
}
