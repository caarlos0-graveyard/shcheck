package sh

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
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

func install(urls map[string]string, path string) (string, error) {
	var url = urls[runtime.GOOS+runtime.GOARCH]
	if url == "" {
		return "", fmt.Errorf("no binary for %s %s", runtime.GOOS, runtime.GOARCH)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return path, download(url, path)
	}
	return path, nil
}
