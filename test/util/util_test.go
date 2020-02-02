package util

import (
	"os/exec"
	"path"
	"syscall"
	"testing"
)

func TestLookPath(t *testing.T) {
	// 寻找 ls 命令的绝对路径
	path, err := exec.LookPath("ls")
	if err != nil {
		t.Error(err)
	}
	t.Logf("ls path: %s \n", path)
}

// 切换运行时目录
func TestChangeRunDir(t *testing.T) {
	err := syscall.Chdir("/root")
	if err != nil {
		t.Error(err)
	}
	cmd := exec.Command("pwd")
	bs, _ := cmd.CombinedOutput()
	t.Log(string(bs))
}

func TestPathJoin(t *testing.T) {
	newPath := path.Join("/root/", "busybox.tar")
	t.Log(newPath)
}
