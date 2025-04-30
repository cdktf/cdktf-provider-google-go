// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionGcsNotebookSource struct {
	// The Cloud Storage uri pointing to the ipynb file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/colab_notebook_execution#uri ColabNotebookExecution#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// The version of the Cloud Storage object to read.
	//
	// If unset, the current version of the object is read. See https://cloud.google.com/storage/docs/metadata#generation-number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/colab_notebook_execution#generation ColabNotebookExecution#generation}
	Generation *string `field:"optional" json:"generation" yaml:"generation"`
}

