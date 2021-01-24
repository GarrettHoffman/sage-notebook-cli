package cmd

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/garretthoffman/sage-notebook-cli/console"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	version       = "0.0.1"
	defaultRegion = "us-east-1"
	runtimeMacOS  = "darwin"
)

var (
	output  ConsoleOutput
	profile string
	region  string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "sage-notebook",
	Short: "Create, configure and manage AWS Sagemaker notebook envi from your command line",
	Long:  "sage notebook is a command-line interface to manage hosted notebook environments using AWS Sagemaker.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		output = ConsoleOutput{}

		if cmd.Parent().Name() == "sage-notebook" {
			return
		}

		if verbose {
			verbose = true
			console.Verbose = true
			output.Verbose = true
		}

		if terminal.IsTerminal(int(os.Stdout.Fd())) {
			console.Color = true
			output.Color = true

			if runtime.GOOS == runtimeMacOS {
				output.Emoji = true
			}
		}

		envAwsDeafultRegion := os.Getenv("AWS_DEFAULT_REGION")
		envAwsRegion := os.Getenv("AWS_REGION")

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

		if err != nil {
			console.ErrorExit(err, "Could not load aws config for profile %s from ~/.aws/config", profile)
		}

		if region != "" {
			cfg.Region = region
		} else if envAwsDeafultRegion != "" {
			cfg.Region = envAwsDeafultRegion
		} else if envAwsRegion != "" {
			cfg.Region = envAwsRegion
		} else if cfg.Region == "" {
			cfg.Region = defaultRegion
		}
	},
}

func Execute() {
	rootCmd.Version = version
	err := rootCmd.Execute()

	if err != nil {
		fmt.Errorf("Could not run command sage notebook %v", err)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().StringVar(&region, "region", "", `AWS region (default "us-east-1")`)
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "", `AWS profile (default "default")`)
}
