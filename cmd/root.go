package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "littlebird",
		Short: "A little birdy CLI for twitter.",
	}

	rootCommand.AddCommand(
		newStartCommand(),
	)

	return rootCommand
}

func Execute() error {
	return newRootCommand().Execute()
}
