// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkload struct {
	// dataflow_flex_template_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_pipeline_pipeline#dataflow_flex_template_request DataPipelinePipeline#dataflow_flex_template_request}
	DataflowFlexTemplateRequest *DataPipelinePipelineWorkloadDataflowFlexTemplateRequest `field:"optional" json:"dataflowFlexTemplateRequest" yaml:"dataflowFlexTemplateRequest"`
	// dataflow_launch_template_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_pipeline_pipeline#dataflow_launch_template_request DataPipelinePipeline#dataflow_launch_template_request}
	DataflowLaunchTemplateRequest *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequest `field:"optional" json:"dataflowLaunchTemplateRequest" yaml:"dataflowLaunchTemplateRequest"`
}

