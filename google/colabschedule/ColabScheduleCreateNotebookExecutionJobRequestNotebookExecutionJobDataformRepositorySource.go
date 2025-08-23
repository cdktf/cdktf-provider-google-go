// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource struct {
	// The resource name of the Dataform Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/colab_schedule#dataform_repository_resource_name ColabSchedule#dataform_repository_resource_name}
	DataformRepositoryResourceName *string `field:"required" json:"dataformRepositoryResourceName" yaml:"dataformRepositoryResourceName"`
	// The commit SHA to read repository with. If unset, the file will be read at HEAD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/colab_schedule#commit_sha ColabSchedule#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
}

