package ecdsa

import (
    "bytes"
    "context"
    "fmt"
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
        fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --file \"test.pem\"", constants.Name())),
        "",
        fmt.Sprintf("  %s", "# Full system path target file"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --file \"/tmp/test.pem\"", constants.Name())),
        "",
        fmt.Sprintf("  %s", "# Generate a directory if it doesn't exist"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --mkdir --file \"./example/test.pem\"", constants.Name())),
        "",
        fmt.Sprintf("  %s", "# Usage without writing to a file"),
        fmt.Sprintf("  %s", fmt.Sprintf("%s go run . ecdsa --dry-run", constants.Name())),
    }, "\n"),
    PreRunE: func(cmd *cobra.Command, args []string) error {
        ctx := cmd.Context()

        logger := slog.With(slog.String("command", cmd.Name()))

        var buffer bytes.Buffer
        if e := system.ECDSA(ctx, &buffer); e != nil {
            e = fmt.Errorf("unable to write to buffer: %w", e)
            return e
        }

        ctx = context.WithValue(ctx, "content", &buffer)

        if !(dryrun) {
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
                if e := os.MkdirAll(dir, 0o755); e != nil {
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

        }

        cmd.SetContext(ctx)

        return nil
    },
    RunE: func(cmd *cobra.Command, args []string) error {
        ctx := cmd.Context()

        content := ctx.Value("content").(*bytes.Buffer)
        if dryrun {
            defer fmt.Fprintf(os.Stdout, "%s", content.String())

            return nil
        }

        path := ctx.Value("path").(string)
        if e := os.WriteFile(path, content.Bytes(), 0o400); e != nil {
            slog.ErrorContext(ctx, "Unable to Write ECDSA Private Key Contents to File", slog.String("error", e.Error()), slog.String("path", path))
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
    flags.BoolVar(&dryrun, "dry-run", false, "output ecdsa private to standard-output without writing to a file")

    Command.MarkFlagsOneRequired("file", "dry-run")
}
