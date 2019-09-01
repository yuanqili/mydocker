// +build linux

package container

import (
	"github.com/Sirupsen/logrus"
	"os"
	"syscall"
)

// The first process in the container.
// It first mounts the procfs, so that we can call `ps` later.
// - MS_NOEXEC: other processes are not allowed in current fs
// - MS_NOSUID: set-user/group-ID are not allowed
// - MS_NODEV:  default flag since Linux 2.4
func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %v", command)

	_ = syscall.Mount("proc", "/proc", "proc",
		syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV, "")

	argv := []string{command}
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}

	return nil
}
