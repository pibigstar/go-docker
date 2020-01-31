package subsystem

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetCgroupPath(t *testing.T) {
	logrus.Infof(findCgroupMountPoint("memory"))
	logrus.Infof(findCgroupMountPoint("cpu"))
	logrus.Infof(findCgroupMountPoint("cpuset"))
}
