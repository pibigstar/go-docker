package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"go-docker/cgroups/subsystem"
	"go-docker/container"
)

// 创建namespace隔离的容器进程
// 启动容器
var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and cgroups limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container args")
		}
		tty := context.Bool("ti")

		res := &subsystem.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		// cmdArray 为容器运行后，执行的第一个命令信息
		// cmdArray[0] 为命令内容, 后面的为命令参数
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		Run(cmdArray, tty, res)
		return nil
	},
}

// 初始化容器内容,挂载proc文件系统，运行用户执行程序
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		return container.RunContainerInitProcess()
	},
}
