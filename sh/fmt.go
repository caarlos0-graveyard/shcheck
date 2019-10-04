package sh

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/pmezard/go-difflib/difflib"
)

type shfmt struct {
}

// Check a file with shfmt
func (s *shfmt) Check(file string) error {
	shfmt, err := s.Install()
	if err != nil {
		return err
	}
	out, err := exec.Command(shfmt, file).CombinedOutput()
	if err != nil {
		return fmt.Errorf("shfmt failed: %v", string(out))
	}
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("shfmt failed: %v", err)
	}
	if string(out) == string(contents) {
		return nil
	}
	diff, err := difflib.GetUnifiedDiffString(
		difflib.UnifiedDiff{
			A:        difflib.SplitLines(string(contents)),
			B:        difflib.SplitLines(string(out)),
			FromFile: "Original",
			ToFile:   "Fixed",
			Context:  3,
		},
	)
	if err != nil {
		return err
	}
	return fmt.Errorf("shfmt failed: file format is wrong, fix it with shfmt -w %s. diff:\n%s", file, diff)
}

// Install shfmt
func (*shfmt) Install() (string, error) {
	return install(
		map[string]string{
			"linuxamd64":  "https://github.com/mvdan/sh/releases/download/v2.6.4/shfmt_v2.6.4_linux_amd64",
			"darwinamd64": "https://github.com/mvdan/sh/releases/download/v2.6.4/shfmt_v2.6.4_darwin_amd64",
		},
		"/tmp/shfmt",
	)
}
