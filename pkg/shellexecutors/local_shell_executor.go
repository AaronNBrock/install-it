package shellexecutors

import (
	"bytes"
	osexec "os/exec"
)

type LocalShellExecutor struct{}

func (LocalShellExecutor) Exec(name string, arg ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := osexec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
