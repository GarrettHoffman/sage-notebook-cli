package cmd

import (
	"context"
	"os"
	"runtime"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/garretthoffman/sage-notebook-cli/console"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	version      = "0.0.1"
	runtimeMacOS = "darwin"
)

var (
	cfg     aws.Config
	output  ConsoleOutput
	profile string
	verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "sage-notebook",
	Short: "Create, configure and manage AWS Sagemaker notebook envi from your command line",
	Long:  "sage is a command-line interface to manage hosted notebook environments using AWS Sagemaker.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		output = ConsoleOutput{}

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

		profileCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))

		if err != nil {
			console.ErrorExit(err, "Could not load aws config for profile %s from ~/.aws/config", profile)
		}

		cfg = profileCfg
	},
}

func Execute() {
	rootCmd.Version = version
	err := rootCmd.Execute()

	if err != nil {
		output.Fatal(err, "Could not run sage command")
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", `AWS profile`)
}
