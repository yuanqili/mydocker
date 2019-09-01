package cgroups

import (
	"github.com/Sirupsen/logrus"
	"github.com/yuanqili/mydocker/cgroups/subsystems"
)

type CgroupManager struct {
	// cgroup's hierarchical path, relative to the root cgroup path
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path:     path,
		Resource: nil,
	}
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subsystemIns := range subsystems.SubsystemsIns {
		if err := subsystemIns.Apply(c.Path, pid); err != nil {
			logrus.Warnf("apply cgroup failed: %v", err)
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subsystemIns := range subsystems.SubsystemsIns {
		if err := subsystemIns.Set(c.Path, res); err != nil {
			logrus.Warnf("set cgroup failed: %v", err)
			return err
		}
	}
	return nil
}

func (c *CgroupManager) Remove() error {
	for _, subsystemIns := range subsystems.SubsystemsIns {
		if err := subsystemIns.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup failed: %v", err)
			return err
		}
	}
	return nil
}
