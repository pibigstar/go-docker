package main

import (
	"github.com/urfave/cli"
)

const usage = `go-docker`

func main() {
	app := cli.NewApp()
	app.Name = "go-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		runCommand,
		initCommand,
	}
}

