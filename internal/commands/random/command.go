package random

import (
    "strings"

    "github.com/spf13/cobra"

    "github.com/x-ethr/ethr-cli/internal/commands/random/token"
)

var Command = &cobra.Command{
	Use:        "random",
	Aliases:    []string{},
	SuggestFor: nil,
	Short:      "Random Data Generator",
	Long:       "A random data generator for various structures.",
	Example: strings.Join([]string{
	}, "\n"),
	PreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	PostRun:           nil,
	CompletionOptions: cobra.CompletionOptions{},
	TraverseChildren:  false,
	Hidden:            false,
	SilenceErrors:     true,
	SilenceUsage:      false,
}

func init() {
    Command.AddCommand(token.Command)
}
