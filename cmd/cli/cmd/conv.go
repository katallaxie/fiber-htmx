package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

func init() {
	ConvertCmd.Flags().StringVarP(&cfg.Convert.File, "file", "f", "", "file to convert")
	ConvertCmd.Flags().StringVarP(&cfg.Convert.Output, "output", "o", "", "output file")
}

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runConvert(cmd.Context())
	},
}

func runConvert(_ context.Context) error {
	return nil
}
