// +build linux

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/yuanqili/mydocker/container"
	"os"
	"strings"
)

func Run(tty bool, cmds []string) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logrus.Errorf("new parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}

	sendInitCommand(cmds, writePipe)
	_ = parent.Wait()
	os.Exit(0)
}

func sendInitCommand(cmds []string, writePipe *os.File) {
	command := strings.Join(cmds, " ")
	_, _ = writePipe.WriteString(command)
	_ = writePipe.Close()
}
