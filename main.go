package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	_ "go-docker/nsenter"
)

const usage = `go-docker`

func main() {
	app := cli.NewApp()
	app.Name = "go-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		runCommand,
		initCommand,
		commitCommand,
		listCommand,
		logCommand,
		execCommand,
		stopCommand,
		removeCommand,
		networkCommand,
	}
	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
