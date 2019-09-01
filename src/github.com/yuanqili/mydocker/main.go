package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

// ./mydocker run -ti ls -l
func main() {

	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = "a simple docker replica"
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}
	app.Before = func(ctx *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
