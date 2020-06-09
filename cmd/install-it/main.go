package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aaronnbrock/install-it/pkg/packagemanagers"
	"github.com/aaronnbrock/install-it/pkg/shellexecutors"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "install-it",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			sh := shellexecutors.LocalShellExecutor{}
			pm := packagemanagers.ApkPackageManager{MyShell: sh}

			packageName := c.Args().Get(0)

			err := pm.Install(packageName)

			if err != nil {
				fmt.Printf("Failed to install '%s'\n", packageName)
			} else {
				fmt.Printf("Sucessfully installed '%s'\n", packageName)
			}

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	// sh := localExecShell{}
	// pm := apkPackageManager{sh}

	// if err != nil {
	// 	fmt.Printf("Failed to install %s", "bash")
	// } else {
	// 	fmt.Printf("Sucessfully installed %s", "bash")
	// }

}
