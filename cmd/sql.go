package cmd

import (
	"github.com/spf13/cobra"

	"github.com/things-go/ormat/cmd/tool"
)

var sqlCmd = &cobra.Command{
	Use:     "sql",
	Short:   "Generate create table sql",
	Example: "ormat sql",
	RunE: func(*cobra.Command, []string) error {
		initConfig()
		tool.ExecuteCreateSQL()
		return nil
	},
}
