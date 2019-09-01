package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/yuanqili/mydocker/container"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "create a container with namespace and cgroups limit",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
	},
	Action: func(ctx *cli.Context) error {
		if len(ctx.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := ctx.Args().Get(0)
		tty := ctx.Bool("ti")
		Run(tty, cmd)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "init container process, run user's process in container",
	Action: func(ctx *cli.Context) error {
		cmd := ctx.Args().Get(0)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
