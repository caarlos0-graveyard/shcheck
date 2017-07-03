package sh

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

const shellcheckPath = "/tmp/shellcheck"

// ShellcheckOptions pass down options to the shellcheck binary
type ShellcheckOptions struct {
	Exclude []string
}

type shellcheck struct {
	options ShellcheckOptions
}

// Check a file with shellcheck
func (s *shellcheck) Check(file string) error {
	bin, err := binaryFor(s, "shellcheck")
	if err != nil {
		return err
	}
	var args = []string{"--external-sources"}
	if len(s.options.Exclude) != 0 {
		args = append(args, "--exclude", strings.Join(s.options.Exclude, ","))
	}
	args = append(args, file)
	out, err := exec.Command(bin, args...).CombinedOutput()
	if err == nil {
		return nil
	}
	return fmt.Errorf("shellcheck failed: %v", string(out))
}

// Install shellcheck
func (*shellcheck) Install() (string, error) {
	if runtime.GOOS == "linux" {
		return shellcheckPath, download(
			"https://github.com/caarlos0/shellcheck-docker/releases/download/v0.4.6/shellcheck",
			shellcheckPath,
		)
	}
	return "", nil
}
