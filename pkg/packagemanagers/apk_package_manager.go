package packagemanagers

import (
	"fmt"

	"github.com/aaronnbrock/install-it/pkg/shellexecutors"
)

type ApkPackageManager struct {
	MyShell shellexecutors.ShellExecutor
}

func (apkPM ApkPackageManager) Install(packageName string) error {
	_, stderr, err := apkPM.MyShell.Exec("apk", "add", packageName)
	if stderr != "" {
		return fmt.Errorf(stderr)
	}
	return err
}
