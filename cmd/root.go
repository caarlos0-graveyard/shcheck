package cmd

import (
	"fmt"
	"os"

	"github.com/caarlos0/sh/sh"
	"github.com/caarlos0/sh/status"
	zglob "github.com/mattn/go-zglob"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sh",
	Short: "sh validates shell files with both shellcheck and shfmt",
	RunE: func(cmd *cobra.Command, args []string) error {
		matches, err := zglob.Glob(`**/*.*sh`)
		if err != nil {
			return err
		}
		var checks = sh.Checkers()
		for _, file := range matches {
			status.Info(file)
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
		}
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
