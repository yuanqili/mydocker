// +build linux

package container

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// The first process in the container.
// It first mounts the procfs, so that we can call `ps` later.
// - MS_NOEXEC: other processes are not allowed in current fs
// - MS_NOSUID: set-user/group-ID are not allowed
// - MS_NODEV:  default flag since Linux 2.4
func RunContainerInitProcess() error {
	cmds := readUserCommand()
	if cmds == nil || len(cmds) == 0 {
		return fmt.Errorf("empty run command")
	}

	_ = syscall.Mount("proc", "/proc", "proc",
		syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV, "")

	// Lookups the abs path via PATH
	exe, err := exec.LookPath(cmds[0])
	if err != nil {
		logrus.Errorf("exec look path error: %v", err)
		return err
	}

	argv := cmds[0:]
	if err := syscall.Exec(exe, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}

	return nil
}

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		logrus.Errorf("init read pipe error: %v", err)
		return nil
	}
	return strings.Split(string(msg), " ")
}
