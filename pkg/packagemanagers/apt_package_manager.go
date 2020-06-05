package packagemanagers

import "fmt"

type aptPackageManager struct {
	myShell shellexecutors
}

func (aptPM aptPackageManager) install(packageName string) error {
	_, stderr, err := aptPM.myShell.exec("sh", "-c", fmt.Sprintf("apt install %s", packageName))
	if stderr != "" {
		return fmt.Errorf(stderr)
	}
	return err
}
