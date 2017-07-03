package sh

import (
	"io"
	"net/http"
	"os"
	"os/exec"
)

// Checker interface
type Checker interface {
	Check(file string) error
	Install() (string, error)
}

// Options provides options to the underline checkers
type Options struct {
	Shellcheck ShellcheckOptions
}

// Checkers all checkers
func Checkers(opts Options) []Checker {
	return []Checker{
		&shellcheck{opts.Shellcheck},
		&shfmt{},
	}
}

func download(url, target string) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer func() { _ = out.Close() }()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	return os.Chmod(target, 0755)
}

func binaryFor(c Checker, name string) (s string, err error) {
	s, err = exec.LookPath(name)
	if err != nil {
		return c.Install()
	}
	return
}
