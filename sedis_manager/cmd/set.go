package cmd

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a data on Sedis",
	RunE:  setHandle,
}

func setHandle(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	rootCmd.AddCommand(setCmd)
}
