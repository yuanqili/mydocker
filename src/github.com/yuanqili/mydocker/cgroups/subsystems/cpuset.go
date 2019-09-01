package subsystems

type CpusetSubsystem struct {
}

func (c CpusetSubsystem) Name() string {
	panic("implement me")
}

func (c CpusetSubsystem) Set(path string, res *ResourceConfig) error {
	panic("implement me")
}

func (c CpusetSubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (c CpusetSubsystem) Remove(path string) error {
	panic("implement me")
}
