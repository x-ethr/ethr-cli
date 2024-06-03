package token

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/x-ethr/ethr-cli/internal/constants"
)

const set = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var Command = &cobra.Command{
	Use:        "token",
	Aliases:    []string{},
	SuggestFor: nil,
	Short:      "Random Token Generator",
	Long:       "The following command generates a random, url-safe and cryptographically secure token.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# General command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s ecdsa --file \"test.pem\"", constants.Name())),
		"",
	}, "\n"),
	PreRunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var result []byte
		for i := 0; i < length; i++ {
			random, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
			if err != nil {
				return err
			}

			result = append(result, set[random.Int64()])
		}

		fmt.Fprintf(os.Stdout, "%s\n", string(result[:length]))

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

	flags.IntVar(&length, "length", 32, "the desired length of the token")
}
