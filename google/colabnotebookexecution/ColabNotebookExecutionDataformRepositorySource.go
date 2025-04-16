// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionDataformRepositorySource struct {
	// The resource name of the Dataform Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/colab_notebook_execution#dataform_repository_resource_name ColabNotebookExecution#dataform_repository_resource_name}
	DataformRepositoryResourceName *string `field:"required" json:"dataformRepositoryResourceName" yaml:"dataformRepositoryResourceName"`
	// The commit SHA to read repository with. If unset, the file will be read at HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/colab_notebook_execution#commit_sha ColabNotebookExecution#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
}

