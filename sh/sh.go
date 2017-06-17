package sh

import (
	"io"
	"net/http"
	"os"
)

// Checker interface
type Checker interface {
	Check(file string) error
	Install() error
}

// Checkers all checkers
func Checkers() []Checker {
	return []Checker{
		&shellcheck{},
		&shfmt{},
	}
}

func download(url, target string) error {
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if _, err = io.Copy(out, resp.Body); err != nil {
		return err
	}
	return os.Chmod(target, 0755)
}
