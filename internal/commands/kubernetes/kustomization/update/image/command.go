package image

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"sigs.k8s.io/kustomize/api/types"

	"github.com/x-ethr/ethr-cli/internal/constants"
	"github.com/x-ethr/ethr-cli/internal/log"
)

var Command = &cobra.Command{
	Use:        "image",
	Aliases:    []string{"img"},
	SuggestFor: nil,
	Short:      "The example's command short-description",
	Long:       "The example's command long-description -- value should be in full sentences, and can span multiple lines.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# General command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s kubernetes kustomization update image --file ./test-data/update-image/kustomization.yaml --image service:latest --name example --tag 1.0.0 --registry private.registry.io", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# With verbose logging"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s kubernetes kustomization update image --verbosity trace --file ./test-data/update-image/kustomization.yaml --image service:latest --name example --tag 1.0.0 --registry private.registry.io", constants.Name())),
		"",
		fmt.Sprintf("  %s", "# Only write content to standard-output (dry-run)"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s kubernetes kustomization update image --file ./test-data/update-image/kustomization.yaml --image service:latest --name example --tag 1.0.0 --registry private.registry.io --dry-run", constants.Name())),
	}, "\n"),
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		logger := slog.With(slog.String("command", cmd.Name()))

		logger.Log(ctx, log.Info, "Parent", slog.String("value", cmd.Parent().Name()))

		cwd, e := os.Getwd()
		if e != nil {
			e = fmt.Errorf("unable to get current working directory: %w", e)
			return e
		}

		path := filepath.Join(cwd, file)

		logger.Log(ctx, log.Debug, "File", slog.String("value", path), slog.String("extension", filepath.Ext(path)))

		if filepath.Ext(path) != ".yml" && filepath.Ext(path) != ".yaml" {
			e = fmt.Errorf("invalid file extension - expecting (\".yml\" | \".yaml\"): %s", filepath.Ext(path))
		} else if _, e := os.Stat(path); errors.Is(e, os.ErrNotExist) {
			e = fmt.Errorf("file does not exist: %s", path)
			return e
		}

		ctx = context.WithValue(ctx, "path", path)

		content, e := os.ReadFile(path)
		if e != nil {
			e = fmt.Errorf("unable to read file: %w", e)
			return e
		}

		var buffer bytes.Buffer

		if size, e := buffer.Write(content); e != nil || size != len(content) {
			e = fmt.Errorf("unable to write to buffer: %w", e)
			return e
		}

		ctx = context.WithValue(ctx, "content", buffer)

		cmd.SetContext(ctx)

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		logger := slog.With(slog.String("command", cmd.Name()))

		logger.Log(ctx, log.Info, "Parent", slog.String("value", cmd.Parent().Name()))

		content, path := ctx.Value("content").(bytes.Buffer), ctx.Value("path").(string)

		var kustomization *types.Kustomization
		if e := yaml.Unmarshal(content.Bytes(), &kustomization); e != nil {
			e = fmt.Errorf("unable to unmarshal kustomization: %w", e)
			return e
		}

		new := image
		if registry != "" {
			new = fmt.Sprintf("%s/%s", registry, name)
		}

		if len(kustomization.Images) == 0 {
			kustomization.Images = []types.Image{
				{
					Name:    image,
					NewName: new,
					NewTag:  tag,
				},
			}
		} else {
			kustomization.Images[0].Name = image
			kustomization.Images[0].NewName = new
			kustomization.Images[0].NewTag = tag
		}

		output, e := yaml.Marshal(kustomization)
		if e != nil {
			e = fmt.Errorf("unable to marshal kustomization: %w", e)
			return e
		}

		if test {
			fmt.Fprintf(os.Stdout, "%s\n", string(output))

			return nil
		}

		return os.WriteFile(path, output, 0o644)
	},
	TraverseChildren: true,
	Hidden:           false,
	SilenceErrors:    true,
	SilenceUsage:     false,
}

func init() {
	flags := Command.Flags()

	flags.StringVar(&file, "file", "", "target file that contains the kubernetes service")
	flags.StringVar(&image, "image", "", "the updated image's name")
	flags.StringVar(&name, "name", "", "the image's reference (e.g. \"service-name:latest\")")
	flags.StringVar(&tag, "tag", "", "the new image's tag")
	flags.StringVar(&registry, "registry", "", "the container registry")

	flags.BoolVar(&test, "dry-run", false, "write updated contents to standard-output instead of file")

	if e := Command.MarkFlagRequired("file"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}

	if e := Command.MarkFlagRequired("image"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}

	if e := Command.MarkFlagRequired("name"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}

	if e := Command.MarkFlagRequired("tag"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}
}
