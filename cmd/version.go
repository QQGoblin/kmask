package cmd

import (
	"fmt"
	"github.com/QQGoblin/kmask/pkg/version"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {

	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%#v\n", version.Get())
		},
	}

}
