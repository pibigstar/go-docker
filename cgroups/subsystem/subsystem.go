package subsystem

// 资源限制配置
type ResourceConfig struct {
	// 内存限制
	MemoryLimit string
	// CPU时间片权重
	CpuShare string
	// CPU核数
	CpuSet string
}

/**
将cgroup抽象成path, 因为在hierarchy中，cgroup便是虚拟的路径地址
*/
type Subystem interface {
	// 返回subsystem名字，如 cpu,memory
	Name() string
	// 设置cgroup在这个subSystem中的资源限制
	Set(cgroupPath string, res *ResourceConfig) error
	// 移除这个cgroup资源限制
	Remove(cgroupPath string) error
	// 将某个进程添加到cgroup中
	Apply(cgroupPath string, pid int) error
}

var (
	Subsystems = []Subystem{
		&MemorySubSystem{},
		&CpuSubSystem{},
		&CpuSetSubSystem{},
	}
)
