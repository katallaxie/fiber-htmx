package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var cfg = DefaultConfig()

const (
	versionFmt = "%s (%s %s)"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(ConvertCmd)

	RootCmd.SilenceErrors = true
	RootCmd.SilenceUsage = true
}

func initConfig() {
}

var RootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runRoot(cmd.Context())
	},
	Version: fmt.Sprintf(versionFmt, version, commit, date),
}

func runRoot(_ context.Context) error {
	return nil
}
