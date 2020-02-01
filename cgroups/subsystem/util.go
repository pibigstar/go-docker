package subsystem

import (
	"bufio"
	"os"
	"path"
	"strings"

	"github.com/sirupsen/logrus"
)

// 获取cgroup在文件系统中的绝对路径
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRootPath, err := findCgroupMountPoint(subsystem)
	if err != nil {
		logrus.Errorf("find cgroup mount point, err: %s", err.Error())
		return "", err
	}
	cgroupTotalPath := path.Join(cgroupRootPath, cgroupPath)
	_, err = os.Stat(cgroupTotalPath)
	if err != nil && os.IsNotExist(err) {
		if err := os.MkdirAll(cgroupTotalPath, 0755); err != nil {
			return "", err
		}
	}

	return cgroupTotalPath, nil
}

// 找到挂载了 subsystem 的hierarchy cgroup根节点所在的目录
func findCgroupMountPoint(subystem string) (string, error) {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subystem && len(fields) > 4 {
				return fields[4], nil
			}
		}
	}
	return "", scanner.Err()
}
