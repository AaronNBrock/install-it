package packagemanagers

import (
	"fmt"
	"os"

	"github.com/aaronnbrock/install-it/pkg/shellexecutors"
)

type AptPackageManager struct {
	MyShell shellexecutors.ShellExecutor
}

func NewAptPackageManager(myShell shellexecutors.ShellExecutor) AptPackageManager {
	aptPM := AptPackageManager{
		MyShell: myShell,
	}
	aptPM.MyShell.Exec("apt-get", "update")
	return aptPM
}

func (aptPM AptPackageManager) Install(packageName string) error {
	_, stderr, err := aptPM.MyShell.Exec("apt-get", "install", "-y", packageName)
	// fmt.Println(stout, stderr, err)
	if stderr != "" {
		fmt.Fprint(os.Stderr, stderr)
	}
	return err
}
