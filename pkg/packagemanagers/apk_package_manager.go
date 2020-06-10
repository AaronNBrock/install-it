package packagemanagers

import (
	"fmt"
	"os"

	"github.com/aaronnbrock/install-it/pkg/shellexecutors"
)

type ApkPackageManager struct {
	MyShell shellexecutors.ShellExecutor
}

func (apkPM ApkPackageManager) Install(packageName string) error {
	_, stderr, err := apkPM.MyShell.Exec("apk", "add", packageName)
	if stderr != "" {
		fmt.Fprint(os.Stderr, stderr)
	}
	return err
}
