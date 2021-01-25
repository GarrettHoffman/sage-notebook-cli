package cmd

import (
	"github.com/pkg/browser"

	"github.com/garretthoffman/sage-notebook-cli/sagemaker"
	"github.com/spf13/cobra"
)

type launchJupyterLabOperation struct {
	sagemaker            sagemaker.Client
	notebookInstanceName string
	output               Output
}

func (o launchJupyterLabOperation) execute() {
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

	jupyterLab := "https://" + notebookInstance.Url + "/lab"
	err = browser.OpenURL(jupyterLab)

	if err != nil {
		o.output.Fatal(err, "Error launching jupyterlab for instance %s", o.notebookInstanceName)
	}
}

var launchJupyterlabCmd = &cobra.Command{
	Use:   "jupyter-lab <notebook-instance-name>",
	Short: "Launch jupyter lab for the notebook instance",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		launchJupyterLabOperation{
			sagemaker:            sagemaker.New(cfg),
			notebookInstanceName: args[0],
			output:               output,
		}.execute()
	},
}

func init() {
	launchCmd.AddCommand(launchJupyterlabCmd)
}
