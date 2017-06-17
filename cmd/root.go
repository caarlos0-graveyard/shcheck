package cmd

import (
	"fmt"
	"os"

	"github.com/caarlos0/shcheck/sh"
	"github.com/caarlos0/shcheck/status"
	zglob "github.com/mattn/go-zglob"
	"github.com/spf13/cobra"
)

var ignores []string

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
		var checks = sh.Checkers()
		for _, file := range files {
			status.Info(file)
			if ignore(ignores, file) {
				fmt.Printf("\n")
				continue
			}
			var errors []error
			for _, check := range checks {
				if err := check.Check(file); err != nil {
					errors = append(errors, err)
				}
			}
			if len(errors) == 0 {
				status.Success(file)
				continue
			}
			status.Fail(file)
			for _, err := range errors {
				fmt.Println(err)
			}
			fmt.Printf("\n\n")
			fail = true
		}

		if fail {
			fmt.Printf("\n\nsome checks failed. check logs above\n")
			os.Exit(1)
		}
	},
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
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
