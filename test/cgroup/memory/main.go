package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const (
	// 挂载了 memory subsystem的hierarchy的根目录位置
	cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"
)

func main() {

	if os.Args[0] == "/proc/self/exe" {
		//容器进程
		fmt.Printf("current pid %d \n", syscall.Getpid())

		cmd := exec.Command("sh", "-c", "stress --vm-bytes 200m --vm-keep -m 1")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err !=nil {
			panic(err)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err:= cmd.Start()
	if err != nil {
		panic(err)
	}
	// 得到 fork出来进程映射在外部命名空间的pid
	fmt.Printf("%+v", cmd.Process.Pid)

	// 创建子cgroup
	newCgroup := path.Join(cgroupMemoryHierarchyMount, "testmemorylimit")
	if err := os.Mkdir(newCgroup, 0755); err !=nil {
		panic(err)
	}
	// 将容器进程放到子cgroup中
	if err := ioutil.WriteFile(path.Join(newCgroup, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err !=nil {
		panic(err)
	}
	// 限制cgroup的内存使用
	if err := ioutil.WriteFile(path.Join(newCgroup, "memory.limit_in_bytes"), []byte("100m"), 0644); err !=nil {
		panic(err)
	}
	cmd.Process.Wait()
}
