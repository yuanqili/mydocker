package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
)

type MemorySubsystem struct {
}

func (s MemorySubsystem) Name() string {
	panic("implement me")
}

func (s MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if res.MemoryLimit == "" {
		return nil
	}

	subsystemCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(subsystemCgroupPath, "memory.limit_in_bytes"),
		[]byte(res.MemoryLimit), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup memory failed, %v", err)
	} else {
		return nil
	}
}

func (s MemorySubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (s MemorySubsystem) Remove(path string) error {
	panic("implement me")
}
