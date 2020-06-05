package packagemanagers

import (
	"fmt"
)

type apkPackageManager struct {
	myShell main.shellexecutors
}

func (apkPM apkPackageManager) install(packageName string) error {
	_, stderr, err := apkPM.myShell.exec("sh", "-c", fmt.Sprintf("apk add %s", packageName))
	if stderr != "" {
		return fmt.Errorf(stderr)
	}
	return err
}
