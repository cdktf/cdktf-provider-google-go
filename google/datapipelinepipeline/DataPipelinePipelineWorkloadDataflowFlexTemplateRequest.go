// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkloadDataflowFlexTemplateRequest struct {
	// launch_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_pipeline_pipeline#launch_parameter DataPipelinePipeline#launch_parameter}
	LaunchParameter *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter `field:"required" json:"launchParameter" yaml:"launchParameter"`
	// The regional endpoint to which to direct the request. For example, us-central1, us-west1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_pipeline_pipeline#location DataPipelinePipeline#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the Cloud Platform project that the job belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_pipeline_pipeline#project_id DataPipelinePipeline#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// If true, the request is validated but not actually executed. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/data_pipeline_pipeline#validate_only DataPipelinePipeline#validate_only}
	ValidateOnly interface{} `field:"optional" json:"validateOnly" yaml:"validateOnly"`
}

