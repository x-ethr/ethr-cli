package ecdsa

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/x-ethr/ethr-cli/internal/constants"
	"github.com/x-ethr/ethr-cli/internal/log"
	"github.com/x-ethr/ethr-cli/internal/system"
)

var Command = &cobra.Command{
	Use:        "ecdsa",
	Aliases:    []string{},
	SuggestFor: nil,
	Short:      "ECDSA PEM Private Key Generator",
	Long:       "The following command generates an ECDSA, PEM private key according to user-provided specification.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# General command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s ecdsa --file \"test.pem\"", constants.Name())),
		// "",
		// fmt.Sprintf("  %s", "# Usage without writing to a file"),
		// fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --dry-run", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Full system path target file"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s ecdsa --file \"/tmp/test.pem\"", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Generate a directory if it doesn't exist"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s ecdsa --mkdir --file \"./example/test.pem\"", constants.Name())),
		// "",
		// fmt.Sprintf("  %s", "# Base64-Encode the key's contents"),
		// fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --b64", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Base64-Encode the key's contents and write to file"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s ecdsa --b64 --file \"base64-encoded-key.pem\"", constants.Name())),
	}, "\n"),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		logger := slog.With(slog.String("command", cmd.Name()))

		var private, public bytes.Buffer
		if e := system.ECDSA(ctx, &private, &public); e != nil {
			e = fmt.Errorf("unable to write to buffer: %w", e)
			return e
		}

		if b64 {
			{
				var content bytes.Buffer
				if _, e := base64.NewEncoder(base64.StdEncoding, &content).Write(private.Bytes()); e != nil {
					return e
				}

				private.Reset()
				if _, e := io.Copy(&private, &content); e != nil {
					return e
				}
			}

			{
				var content bytes.Buffer
				if _, e := base64.NewEncoder(base64.StdEncoding, &content).Write(public.Bytes()); e != nil {
					return e
				}

				public.Reset()
				if _, e := io.Copy(&public, &content); e != nil {
					return e
				}
			}
		}

		ctx = context.WithValue(ctx, "private", &private)
		ctx = context.WithValue(ctx, "public", &public)

		cwd, e := os.Getwd()
		if e != nil {
			e = fmt.Errorf("unable to get current working directory: %w", e)
			return e
		}

		path := filepath.Join(cwd, file)
		if filepath.IsAbs(file) {
			path = file
		}

		dir := filepath.Dir(path)
		if !(system.Exists(dir)) && !(mkdir) {
			return fmt.Errorf("directory does not exist: %s", dir)
		} else if !(system.Exists(dir)) {
			if e := os.MkdirAll(dir, 0o775); e != nil {
				return e
			}
		}

		extension := filepath.Ext(path)
		if extension == "" {
			path = fmt.Sprintf("%s.pem", file)
			extension = filepath.Ext(path)
		} else if extension != ".pem" {
			e = fmt.Errorf("unsupported file extension (\"%s\"): must be *.pem or unspecified", extension)
		}

		ctx = context.WithValue(ctx, "path", path)

		logger.Log(ctx, log.Debug, "File", slog.String("value", path), slog.String("extension", extension))

		cmd.SetContext(ctx)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		private, public := ctx.Value("private").(*bytes.Buffer), ctx.Value("public").(*bytes.Buffer)

		path := ctx.Value("path").(string)
		partial := strings.TrimSuffix(path, ".pem")
		privatepath := fmt.Sprintf("%s.private.pem", partial)
		publicpath := fmt.Sprintf("%s.public.pem", partial)
		if e := os.WriteFile(privatepath, private.Bytes(), 0o664); e != nil {
			slog.ErrorContext(ctx, "Unable to Write ECDSA Private Key Contents to File", slog.String("error", e.Error()), slog.String("path", path))
			return e
		}

		if e := os.WriteFile(publicpath, public.Bytes(), 0o664); e != nil {
			slog.ErrorContext(ctx, "Unable to Write ECDSA Public Key Contents to File", slog.String("error", e.Error()), slog.String("path", path))
			return e
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

	flags.StringVar(&file, "file", "", "a relative or full-system path to the target generated ecdsa, pem file")
	flags.BoolVar(&mkdir, "mkdir", false, "generate the target directory if it does not exist")
	flags.BoolVar(&b64, "b64", false, "base64-encode the key's contents - useful when working with kubernetes or secrets")
	// flags.BoolVar(&dryrun, "dry-run", false, "output ecdsa private to standard-output without writing to a file")

	if e := Command.MarkFlagRequired("file"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}
}
