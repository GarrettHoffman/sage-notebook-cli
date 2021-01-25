package cmd

import (
	"github.com/pkg/browser"

	"github.com/garretthoffman/sage-notebook-cli/sagemaker"
	"github.com/spf13/cobra"
)

type launchJupyterOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
	output               Output
}

func (o launchJupyterOperation) execute() {
	o.output.Debug("Describing Notebook Instance: %s [API=sagemaker Action=DescribeNotebookInstance]", o.notebookInstanceName)
	notebookInstance, err := o.sagemaker.DescribeNotebookInstance(o.notebookInstanceName)

	if err != nil {
		o.output.Fatal(err, "No notebook instance %s", o.notebookInstanceName)
		return
	}

	if notebookInstance.NotebookInstanceStatus != "InService" {
		o.output.Info("Notebook instance status must be InService to launch Jupyter, run sage notebook start %s", o.notebookInstanceName)
		return
	}

	jupyter := "https://" + notebookInstance.Url + "/"
	err = browser.OpenURL(jupyter)

	if err != nil {
		o.output.Fatal(err, "Error launching jupyter for instance %s", o.notebookInstanceName)
	}
}

var launchJupyterCmd = &cobra.Command{
	Use:   "jupyter <notebook-instance-name>",
	Short: "Launch jupyter  for the notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		launchJupyterOperation{
			sagemaker:            sagemaker.New(cfg),
			notebookInstanceName: args[0],
			output:               output,
		}.execute()
	},
}

func init() {
	launchCmd.AddCommand(launchJupyterCmd)
}
