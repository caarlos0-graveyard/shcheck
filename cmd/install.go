package cmd

import (
	"github.com/caarlos0/shcheck/sh"
	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "quickly install the binaries to /tmp to use them later",
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, check := range sh.Checkers() {
			if err := check.Install(); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(installCmd)
}
