package sagemaker

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsSagemaker "github.com/aws/aws-sdk-go-v2/service/sagemaker"
)

// Client represents a method for accessing Sagemaker
type Client interface {
	ListNotebookInstances() (NotebookInstances, error)
	DescribeNotebookInstance(string) (NotebookInstance, error)
	StopNotebookInstance(string) error
}

// SDKClient implements access to Sagemaker via the AWS SDK
type SDKClient struct {
	client *awsSagemaker.Client
}

// New returns a SDKClient configured with a given configuration
func New(cfg aws.Config) SDKClient {
	return SDKClient{
		client: awsSagemaker.NewFromConfig(cfg),
	}
}
