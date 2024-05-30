package kubernetes

import (
	"github.com/spf13/cobra"

	"github.com/x-ethr/ethr-cli/internal/commands/kubernetes/kustomization"
)

var Command = &cobra.Command{
	Use:                    "kubernetes",
	Short:                  "K8s API Command(s)",
	Long:                   "Kubernetes related commands for abstracting away API complexities or simplifying CLI-related usage.",
	Aliases:                []string{},
	SuggestFor:             nil,
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
	Command.AddCommand(kustomization.Command)
}
