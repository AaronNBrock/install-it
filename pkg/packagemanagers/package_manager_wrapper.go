package packagemanagers

import (
	"fmt"
)

type PackageManagerWrapper struct {
	PackageManagers []PackageManager
}

func (pmw PackageManagerWrapper) Install(packageName string) error {
	var err error

	for _, pm := range pmw.PackageManagers {
		err = pm.Install(packageName)
		if err == nil {
			break
		} else {
			fmt.Println(err)
		}
	}

	if err != nil {
		return fmt.Errorf("unable to install %s", packageName)
	}

	return err
}
