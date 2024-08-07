package cmd

import (
	"github.com/QQGoblin/kmask/pkg/core"
	"github.com/spf13/cobra"
)

var (
	config       string
	picture      string
	impassphrase string
)

func init() {
	ImportCommand.PersistentFlags().StringVarP(&config, "config", "c", core.DefaultExportConfigName, "secret config file")
	ImportCommand.PersistentFlags().StringVarP(&picture, "picture", "b", "", "background picture for store secret")
	ImportCommand.PersistentFlags().StringVarP(&impassphrase, "passphrase", "p", "", "passphrase for secret database file")
}

var ImportCommand = &cobra.Command{
	Use:   "import",
	Short: "import secret from json",
	RunE: func(cmd *cobra.Command, args []string) error {
		return core.Import(config, impassphrase, picture)
	},
}
