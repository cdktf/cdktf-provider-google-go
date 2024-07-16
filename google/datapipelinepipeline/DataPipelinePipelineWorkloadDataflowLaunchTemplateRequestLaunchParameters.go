// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParameters struct {
	// The job name to use for the created job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_pipeline_pipeline#job_name DataPipelinePipeline#job_name}
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_pipeline_pipeline#environment DataPipelinePipeline#environment}
	Environment *DataPipelinePipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// The runtime parameters to pass to the job.
	//
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_pipeline_pipeline#parameters DataPipelinePipeline#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Map of transform name prefixes of the job to be replaced to the corresponding name prefixes of the new job.
	//
	// Only applicable when updating a pipeline.
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_pipeline_pipeline#transform_name_mapping DataPipelinePipeline#transform_name_mapping}
	TransformNameMapping *map[string]*string `field:"optional" json:"transformNameMapping" yaml:"transformNameMapping"`
	// If set, replace the existing pipeline with the name specified by jobName with this pipeline, preserving state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_pipeline_pipeline#update DataPipelinePipeline#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

