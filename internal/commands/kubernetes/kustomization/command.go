package kustomization

import (
    "fmt"
    "strings"

    "github.com/spf13/cobra"

    "github.com/x-ethr/ethr-cli/internal/commands/kubernetes/kustomization/update"

    "github.com/x-ethr/ethr-cli/internal/constants"
)

var Command = &cobra.Command{
    Use:        "kustomization",
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
    SilenceErrors:          true,
    TraverseChildren:       true,
}

func init() {
    Command.AddCommand(update.Command)
}
