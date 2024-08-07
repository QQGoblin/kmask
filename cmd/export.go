package cmd

import (
	"github.com/QQGoblin/kmask/pkg/core"
	"github.com/spf13/cobra"
)

var (
	data         string
	expassphrase string
	all          bool
	output       string
)

func init() {
	ExportCommand.PersistentFlags().StringVarP(&data, "database", "d", core.DefaultDBName, "secret database file")
	ExportCommand.PersistentFlags().StringVarP(&expassphrase, "passphrase", "p", "", "passphrase for secret database file")
	ExportCommand.PersistentFlags().StringVarP(&output, "output", "o", "", "output directory")
	ExportCommand.PersistentFlags().BoolVar(&all, "all", false, "export all secret to json file")
}

var ExportCommand = &cobra.Command{
	Use:   "export",
	Short: "export secret from database",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Export(data, expassphrase, output, all)
	},
}
