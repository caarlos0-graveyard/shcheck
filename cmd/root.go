package cmd

import (
	"fmt"
	"os"

	"github.com/caarlos0/shcheck/sh"
	"github.com/caarlos0/shcheck/status"
	zglob "github.com/mattn/go-zglob"
	"github.com/spf13/cobra"
)

var (
	ignores            []string
	shellcheckExcludes []string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sh",
	Short: "sh validates shell files with both shellcheck and shfmt",
	Run: func(cmd *cobra.Command, args []string) {
		var fail bool
		files, err := zglob.Glob(`**/*.*sh`)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, file := range files {
			if err := check(file); err != nil {
				fail = true
			}
		}
		if fail {
			fmt.Printf("\nsome checks failed. check output above.\n")
			os.Exit(1)
		}
	},
}

func check(file string) error {
	if ignore(ignores, file) {
		status.Ignore(file)
		return nil
	}
	var options = sh.Options{
		Shellcheck: sh.ShellcheckOptions{
			Exclude: shellcheckExcludes,
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

func init() {
	RootCmd.PersistentFlags().StringSliceVar(
		&ignores,
		"ignore",
		[]string{},
		"ignore specific folder of file patterns",
	)
	RootCmd.PersistentFlags().StringSliceVar(
		&shellcheckExcludes,
		"shellcheck-exclude",
		[]string{},
		"pass arguments to shellcheck --exclude option",
	)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
