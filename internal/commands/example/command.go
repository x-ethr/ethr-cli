package example

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/cobra"

	"github.com/x-ethr/ethr-cli/internal/constants"
	"github.com/x-ethr/ethr-cli/internal/log"
	"github.com/x-ethr/ethr-cli/internal/marshalers"
	"github.com/x-ethr/ethr-cli/internal/types/output"
)

var Command = &cobra.Command{
	Use:        "example",
	Aliases:    []string{},
	SuggestFor: nil,
	Short:      "The example's command short-description",
	Long:       "The example's command long-description -- value should be in full sentences, and can span multiple lines.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# General command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\"", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Extended usage demonstrating configuration of default(s)"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\" --output json", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Display help information and command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --help", constants.Name())),
	}, "\n"),
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "",
	PreRun:                 func(cmd *cobra.Command, args []string) {},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		slog.Log(ctx, log.Debug, "Example Log Message", slog.Group("command",
			slog.String("name", cmd.Name()),
			slog.Group("flags",
				slog.String("name", name),
				slog.String("output", format.String()),
			),
		))

		var datum = map[string]string{
			"name":   name,
			"output": format.String(),
		}

		switch format {
		case output.JSON:
			buffer, e := marshalers.JSON(datum)
			if e != nil {
				return fmt.Errorf("unable to marshal structure to json: %w", e)
			}

			fmt.Printf("%s", string(buffer))
		case output.YAML:
			buffer, e := marshalers.YAML(datum)
			if e != nil {
				return fmt.Errorf("unable to marshal structure to yaml: %w", e)
			}

			fmt.Printf("%s", string(buffer))
		}

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
	flags := Command.Flags()

	flags.StringVarP(&name, "name", "n", "", "a required example named-string-flag")
	flags.VarP(&format, "output", "o", "structured data format")
	if e := Command.MarkFlagRequired("name"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}
}
