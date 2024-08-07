package main

import (
	"github.com/QQGoblin/kmask/cmd"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
	"os"
)

func NewCommand() *cobra.Command {

	rootCMD := &cobra.Command{
		Use:   "kmask",
		Short: "hide the secret key in picture(png/jpg) through the steganography algorithm(LSB)",
	}

	rootCMD.AddCommand(cmd.VersionCommand())
	rootCMD.AddCommand(cmd.ImportCommand)
	rootCMD.AddCommand(cmd.ExportCommand)
	return rootCMD
}

func main() {

	rootCMD := NewCommand()
	if err := rootCMD.Execute(); err != nil {
		klog.ErrorS(err, "exit with error")
		os.Exit(1)
	}
}
