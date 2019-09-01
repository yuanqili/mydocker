package subsystems

type CpuSubsystem struct {
}

func (c CpuSubsystem) Name() string {
	panic("implement me")
}

func (c CpuSubsystem) Set(path string, res *ResourceConfig) error {
	panic("implement me")
}

func (c CpuSubsystem) Apply(path string, pid int) error {
	panic("implement me")
}

func (c CpuSubsystem) Remove(path string) error {
	panic("implement me")
}
