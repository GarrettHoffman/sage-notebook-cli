package sagemaker

import "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"

func acceleratorTypesToStrings(ats []types.NotebookInstanceAcceleratorType) []string {
	var result []string
	for _, at := range ats {
		result = append(result, string(at))
	}

	return result
}
