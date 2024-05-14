// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest struct {
	// The ID of the Cloud Platform project that the job belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_pipeline_pipeline#project_id DataPipelinePipeline#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// A Cloud Storage path to the template from which to create the job.
	//
	// Must be a valid Cloud Storage URL, beginning with 'gs://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_pipeline_pipeline#gcs_path DataPipelinePipeline#gcs_path}
	GcsPath *string `field:"optional" json:"gcsPath" yaml:"gcsPath"`
	// launch_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_pipeline_pipeline#launch_parameters DataPipelinePipeline#launch_parameters}
	LaunchParameters *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters `field:"optional" json:"launchParameters" yaml:"launchParameters"`
	// The regional endpoint to which to direct the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_pipeline_pipeline#location DataPipelinePipeline#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_pipeline_pipeline#validate_only DataPipelinePipeline#validate_only}.
	ValidateOnly interface{} `field:"optional" json:"validateOnly" yaml:"validateOnly"`
}

