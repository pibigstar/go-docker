package container

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-docker/common"
	"os/exec"
	"path"
)

func CommitContainer(imageName, imagePath string) error {
	if imagePath == "" {
		imagePath = common.RootPath
	}
	imageTar := path.Join(imagePath, fmt.Sprintf("%s.tar", imageName))
	if _, err := exec.Command("tar", "-czf", imageTar, "-C", common.MntPath, ".").CombinedOutput(); err != nil {
		logrus.Errorf("tar container image, file name: %s, err: %v", imageTar, err)
		return err
	}
	return nil
}
