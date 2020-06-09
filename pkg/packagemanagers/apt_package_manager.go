package packagemanagers

import (
	"fmt"

	"github.com/aaronnbrock/install-it/pkg/shellexecutors"
)

type AptPackageManager struct {
	MyShell shellexecutors.ShellExecutor
}

func (aptPM AptPackageManager) Install(packageName string) error {
	_, stderr, err := aptPM.MyShell.Exec("apt", "install", packageName)
	if stderr != "" {
		return fmt.Errorf(stderr)
	}
	return err
}
