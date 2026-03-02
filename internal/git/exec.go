package git

import "os/exec"

type ExecRunner struct{}

func (r *ExecRunner) Run(dir, name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd.Output()
}
