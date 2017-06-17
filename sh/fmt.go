package sh

import "os/exec"
import "fmt"
import "io/ioutil"
import "runtime"

type shfmt struct {
}

const shfmtPath = "/tmp/shfmt"

// Check a file with shfmt
func (s *shfmt) Check(file string) error {
	shfmt, err := binaryFor(s, "shfmt")
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
	if string(contents) != string(out) {
		return fmt.Errorf("shfmt failed: file format is wrong, fix it with shfmt -w %v", file)
	}
	return nil
}

// Install shfmt
func (*shfmt) Install() (string, error) {
	if runtime.GOOS != "linux" {
		return shfmtPath, download(
			"https://github.com/mvdan/sh/releases/download/v1.3.1/shfmt_v1.3.1_linux_amd64",
			shfmtPath,
		)
	}
	return "", fmt.Errorf("platform not supported: %v", runtime.GOOS)
}
