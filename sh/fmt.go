package sh

import "os/exec"
import "fmt"
import "io/ioutil"

type shfmt struct {
}

// Check a file with shfmt
func (*shfmt) Check(file string) error {
	out, err := exec.Command("shfmt", file).CombinedOutput()
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
func (*shfmt) Install() error {
	return nil
}
