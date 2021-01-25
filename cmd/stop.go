package cmd

import (
	"time"

	"github.com/garretthoffman/sage-notebook-cli/sagemaker"
	"github.com/spf13/cobra"
)

type stopOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
	output               Output
}

func (o stopOperation) execute() {
	o.output.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		o.output.Fatal(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "InService" {
		o.output.Info("Notebook %s is not currently in service", o.notebookInstanceName)
		return
	}

	o.output.Debug("Describing Notebook Instance: %s [API=sagemaker Action=StopNotebookInstance]", o.notebookInstanceName)
	err = o.sagemaker.StopNotebookInstance(o.notebookInstanceName)

	if err != nil {
		o.output.Fatal(err, "Error stopping notebook instance %s", o.notebookInstanceName)
		return
	}

	o.output.Info("Stopping notebook instance %s", o.notebookInstanceName)

	notebookStatus := "Stopping"
	for notebookStatus != "Stopped" {
		time.Sleep(5000000000)
		print(".")

		o.output.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
		notebookInstance, err = o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

		if err != nil {
			o.output.Fatal(err, "Error fetching notebook instance status")
			return
		}

		notebookStatus = notebookInstance.NotebookInstanceStatus
	}

	print("\n")
	o.output.Info("Notebook instance %s stopped", o.notebookInstanceName)
}

var stopCmd = &cobra.Command{
	Use:   "stop <notebook-instance-name>",
	Short: "Stop notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stopOperation{
			sagemaker:            sagemaker.New(cfg),
			notebookInstanceName: args[0],
			output:               output,
		}.execute()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
