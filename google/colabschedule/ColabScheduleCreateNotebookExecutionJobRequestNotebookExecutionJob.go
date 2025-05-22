// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob struct {
	// Required. The display name of the Notebook Execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#display_name ColabSchedule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Cloud Storage location to upload the result to. Format:'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#gcs_output_uri ColabSchedule#gcs_output_uri}
	GcsOutputUri *string `field:"required" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#notebook_runtime_template_resource_name ColabSchedule#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"required" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#dataform_repository_source ColabSchedule#dataform_repository_source}
	DataformRepositorySource *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// A duration in seconds with up to nine fractional digits, ending with "s". Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#execution_timeout ColabSchedule#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#execution_user ColabSchedule#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#gcs_notebook_source ColabSchedule#gcs_notebook_source}
	GcsNotebookSource *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/colab_schedule#service_account ColabSchedule#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

