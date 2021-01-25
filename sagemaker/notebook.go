package sagemaker

import (
	"context"
	"time"

	awsSagemaker "github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/garretthoffman/sage-notebook-cli/console"
	"github.com/garretthoffman/sage-notebook-cli/util"
)

// Notebook represents a Sagemaker notebook instance
type NotebookInstance struct {
	AcceleratorTypes                    []string
	AdditionalCodeRepositories          []string
	CreationTime                        time.Time
	DefaultCodeRepository               string
	DirectInternetAccess                string
	FailureReason                       string
	InstanceType                        string
	KmsKeyId                            string
	LastModifiedTime                    time.Time
	NetworkInterfaceId                  string
	NotebookInstanceArn                 string
	NotebookInstanceLifecycleConfigName string
	NotebookInstanceName                string
	NotebookInstanceStatus              string
	RoleArn                             string
	RootAccess                          string
	SecurityGroups                      []string
	SubnetId                            string
	Url                                 string
	VolumeSizeInGB                      int32
}

// Notebooks is a collection of Sagemaker notebook instances.
type NotebookInstances []NotebookInstance

func (sagemaker SDKClient) ListNotebookInstances() (NotebookInstances, error) {
	return sagemaker.listNotebookInstances(&awsSagemaker.ListNotebookInstancesInput{})
}

func (sagemaker SDKClient) DescribeNotebookInstance(name string) (NotebookInstance, error) {
	return sagemaker.describeNotebookInstance(
		&awsSagemaker.DescribeNotebookInstanceInput{
			NotebookInstanceName: &name,
		},
	)
}

func (sagemaker SDKClient) listNotebookInstances(i *awsSagemaker.ListNotebookInstancesInput) (NotebookInstances, error) {
	var notebookInstances NotebookInstances

	for {
		o, err := sagemaker.client.ListNotebookInstances(context.TODO(), i)
		if err != nil {
			return nil, err
		}

		for _, notebookInstance := range o.NotebookInstances {
			notebookInstances = append(notebookInstances,
				NotebookInstance{
					AdditionalCodeRepositories:          notebookInstance.AdditionalCodeRepositories,
					CreationTime:                        *notebookInstance.CreationTime,
					DefaultCodeRepository:               util.DerefOptionalStringPtr(notebookInstance.DefaultCodeRepository),
					InstanceType:                        string(notebookInstance.InstanceType),
					LastModifiedTime:                    *notebookInstance.LastModifiedTime,
					NotebookInstanceArn:                 *notebookInstance.NotebookInstanceArn,
					NotebookInstanceLifecycleConfigName: util.DerefOptionalStringPtr(notebookInstance.NotebookInstanceLifecycleConfigName),
					NotebookInstanceName:                *notebookInstance.NotebookInstanceName,
					NotebookInstanceStatus:              string(notebookInstance.NotebookInstanceStatus),
					Url:                                 *notebookInstance.Url,
				},
			)
		}

		if o.NextToken == nil {
			break
		}

		*i.NextToken = *o.NextToken
	}

	return notebookInstances, nil
}

func (sagemaker SDKClient) describeNotebookInstance(i *awsSagemaker.DescribeNotebookInstanceInput) (NotebookInstance, error) {
	o, err := sagemaker.client.DescribeNotebookInstance(context.TODO(), i)

	console.Debug("%+v", o)

	notebookInstance := NotebookInstance{
		AcceleratorTypes:                    acceleratorTypesToStrings(o.AcceleratorTypes),
		AdditionalCodeRepositories:          o.AdditionalCodeRepositories,
		CreationTime:                        *o.CreationTime,
		DefaultCodeRepository:               util.DerefOptionalStringPtr(o.DefaultCodeRepository),
		DirectInternetAccess:                string(o.DirectInternetAccess),
		FailureReason:                       util.DerefOptionalStringPtr(o.FailureReason),
		InstanceType:                        string(o.InstanceType),
		KmsKeyId:                            util.DerefOptionalStringPtr(o.KmsKeyId),
		LastModifiedTime:                    *o.LastModifiedTime,
		NetworkInterfaceId:                  util.DerefOptionalStringPtr(o.NetworkInterfaceId),
		NotebookInstanceArn:                 *o.NotebookInstanceArn,
		NotebookInstanceLifecycleConfigName: util.DerefOptionalStringPtr(o.NotebookInstanceLifecycleConfigName),
		NotebookInstanceName:                *o.NotebookInstanceName,
		NotebookInstanceStatus:              string(o.NotebookInstanceStatus),
		RoleArn:                             *o.RoleArn,
		RootAccess:                          string(o.RootAccess),
		SecurityGroups:                      o.SecurityGroups,
		SubnetId:                            util.DerefOptionalStringPtr(o.SubnetId),
		Url:                                 *o.Url,
		VolumeSizeInGB:                      *o.VolumeSizeInGB,
	}

	return notebookInstance, err
}
