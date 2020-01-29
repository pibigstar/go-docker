package main

import "go-docker/container"

func Run(cmd string, tty bool) error {
	parent := container.NewParentProcess(cmd, tty)
	if err := parent.Start(); err !=nil {
		return err
	}
	return parent.Wait()
}