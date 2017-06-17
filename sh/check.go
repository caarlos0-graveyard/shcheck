package sh

import "os/exec"
import "fmt"

type shellcheck struct {
}

// Check a file with shellcheck
func (*shellcheck) Check(file string) error {
	out, err := exec.Command("shellcheck", "-x", file).CombinedOutput()
	if err == nil {
		return nil
	}
	return fmt.Errorf("shellcheck failed: %v", string(out))
}

// Install shellcheck
func (*shellcheck) Install() error {
	return nil
}
