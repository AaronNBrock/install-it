package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/aaronnbrock/pkg/shellexecutors"
)

func main() {
	app := &cli.App{
		Name:  "install-it",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			sh := shellexecutors.localExecShell{}
			sh.exec("ls")

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
