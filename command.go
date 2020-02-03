package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go-docker/common"
	"os"

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
		cli.StringFlag{
			Name:  "v",
			Usage: "docker volume",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container args")
		}

		res := &subsystem.ResourceConfig{
			MemoryLimit: ctx.String("m"),
			CpuSet:      ctx.String("cpuset"),
			CpuShare:    ctx.String("cpushare"),
		}
		// cmdArray 为容器运行后，执行的第一个命令信息
		// cmdArray[0] 为命令内容, 后面的为命令参数
		var cmdArray []string
		for _, arg := range ctx.Args() {
			cmdArray = append(cmdArray, arg)
		}

		tty := ctx.Bool("ti")
		volume := ctx.String("v")
		detach := ctx.Bool("d")

		if tty && detach {
			return fmt.Errorf("ti and d paramter can not both provided")
		}

		containerName := ctx.String("name")
		Run(cmdArray, tty, res, volume, containerName)
		return nil
	},
}

// 初始化容器内容,挂载proc文件系统，运行用户执行程序
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process in container. Do not call it outside",
	Action: func(ctx *cli.Context) error {
		logrus.Infof("init come on")
		return container.RunContainerInitProcess()
	},
}

// 导出容器内容
var commitCommand = cli.Command{
	Name:  "commit",
	Usage: "docker commit a container into image",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "c",
			Usage: "export image path",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container name")
		}
		imageName := ctx.Args().Get(0)
		imagePath := ctx.String("c")
		return container.CommitContainer(imageName, imagePath)
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all container",
	Action: func(ctx *cli.Context) error {
		container.ListContainerInfo()
		return nil
	},
}

var logCommand = cli.Command{
	Name:  "logs",
	Usage: "look container log",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container name")
		}
		containerName := ctx.Args().Get(0)
		container.LookContainerName(containerName)
		return nil
	},
}

var execCommand = cli.Command{
	Name:  "exec",
	Usage: "exec a command into container",
	Action: func(ctx *cli.Context) error {
		// 如果环境变量里面有 PID,那么则什么都不执行
		pid := os.Getenv(common.EnvExecPid)
		if pid != "" {
			logrus.Infof("pid callback pid %s, gid: %d", pid, os.Getgid())
			return nil
		}
		if len(ctx.Args()) < 2 {
			return fmt.Errorf("missing container name or command")
		}

		var cmdArray []string
		for _, arg := range ctx.Args().Tail() {
			cmdArray = append(cmdArray, arg)
		}

		containerName := ctx.Args().Get(0)
		container.ExecContainer(containerName, cmdArray)
		return nil
	},
}

var stopCommand = cli.Command{
	Name:  "stop",
	Usage: "stop a container",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing stop container name")
		}
		containerName := ctx.Args().Get(0)
		container.StopContainer(containerName)
		return nil
	},
}

var removeCommand = cli.Command{
	Name:  "rm",
	Usage: "rm a container",
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing remove container name")
		}
		containerName := ctx.Args().Get(0)
		container.RemoveContainer(containerName)
		return nil
	},
}
