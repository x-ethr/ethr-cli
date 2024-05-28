package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/x-ethr/color"

	"github.com/x-ethr/ethr-cli/internal/commands/kubernetes"
)

// Execute runs the root command and handles any CLI execution exception. Additionally,
// all child command(s) are added to the root command.
func Execute(root *cobra.Command) {
	// root.AddCommand(example.Command)

	root.AddCommand(kubernetes.Command)

	if e := root.Execute(); e != nil {
		color.Color().Bold(
			color.Color().Red("error"),
		).Default("-").Italic(
			color.Color().White(e.Error()),
		).Write(os.Stdout)
	}
}
