package common

const (
	RootPath   = "/root/"
	MntPath    = "/root/mnt/"
	BusyBox    = "busybox"
	BusyBoxTar = "busybox.tar"
	WriteLayer = "writeLayer"
)

const (
	Running = "running"
	Stop    = "stopped"
	Exit    = "exited"
)

const (
	DefaultContainerInfoPath = "/var/run/go-docker/"
	ContainerInfoFileName    = "config.json"
	ContainerLogFileName     = "container.log"
)

const (
	EnvExecPid = "docker_pid"
	EnvExecCmd = "docker_cmd"
)
