// +build linux

package subsystems

import (
	"fmt"
	"os"
	"path"
	"strings"
)

// type cgroup string
// type subsystem string

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type Subsystem interface {
	// Returns the name of the subsystem
	Name() string
	// Sets resource limit in a cgroup
	Set(path string, res *ResourceConfig) error
	// Adds a process into a cgroup
	Apply(path string, pid int) error
	// Removes a cgroup
	Remove(path string) error
}

var (
	SubsystemsIns = []Subsystem{
		&CpuSubsystem{},
		&CpusetSubsystem{},
		&MemorySubsystem{},
	}
)

// Gets absolute cgroup path in the filesystem
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)

	_, err := os.Stat(path.Join(cgroupRoot, cgroupPath))
	if err != nil && !(autoCreate && os.IsNotExist(err)) {
		return "", fmt.Errorf("cgroup path error %v", err)
	}

	if os.IsNotExist(err) {
		err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755)
		if err != nil {
			return "", fmt.Errorf("error create cgroup %v", err)
		}
	}

	return path.Join(cgroupRoot, cgroupPath), nil
}

// Finds the CGroup mount point of some subsystem via /proc/self/mountinfo
func FindCgroupMountpoint(subsystem string) string {
	mountinfos, _ := GetMountinfo("/proc/self/mountinfo")
	for _, mountinfo := range mountinfos {
		for _, opt := range strings.Split(mountinfo.SuperOptions, ",") {
			if opt == subsystem {
				return mountinfo.MountPoint
			}
		}
	}

	return ""
}
