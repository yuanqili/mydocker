// +build linux

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/yuanqili/mydocker/container"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	_ = parent.Wait()
	os.Exit(1)
}
