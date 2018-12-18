package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/caarlos0/shcheck/sh"
	"github.com/caarlos0/shcheck/status"
	zglob "github.com/mattn/go-zglob"
)

// nolint: gochecknoglobals,lll
var (
	version            = "master"
	app                = kingpin.New("shcheck", "shcheck validates your scripts against both shellcheck and shfmt")
	ignoredFiles       = app.Flag("ignore", "ignore files or folders").HintOptions("folder/**/*", "*.bash").Short('i').Strings()
	shellcheckExcludes = app.Flag("shellcheck-exclude", "exclude some shellcheck checks").HintOptions("SC1090", "SC1004").Short('e').Strings()
)

func main() {
	app.Version("shfmt version " + version)
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// TODO: also look for executables with a valid shell shebang
	files, err := zglob.Glob(`**/*.*sh`)
	kingpin.FatalIfError(err, "fail to find all shell files")

	var fail bool
	for _, file := range files {
		if err := check(file); err != nil {
			fail = true
		}
	}
	if fail {
		kingpin.Fatalf("\nsome checks failed. check output above.\n")
	}
}

func check(file string) error {
	if ignore(*ignoredFiles, file) {
		status.Ignore(file)
		return nil
	}
	var options = sh.Options{
		Shellcheck: sh.ShellcheckOptions{
			Exclude: *shellcheckExcludes,
		},
	}
	var errors []error
	for _, check := range sh.Checkers(options) {
		if err := check.Check(file); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) == 0 {
		status.Success(file)
		return nil
	}
	status.Fail(file)
	for _, err := range errors {
		fmt.Println(err)
	}
	return fmt.Errorf("check failed")
}

func ignore(patterns []string, file string) bool {
	for _, pattern := range patterns {
		if ok, err := zglob.Match(pattern, file); ok && err == nil {
			return true
		}
	}
	return false
}
