// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameter struct {
	// The job name to use for the created job.
	//
	// For an update job request, the job name should be the same as the existing running job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#job_name DataPipelinePipeline#job_name}
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Cloud Storage path to a file with a JSON-serialized ContainerSpec as content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#container_spec_gcs_path DataPipelinePipeline#container_spec_gcs_path}
	ContainerSpecGcsPath *string `field:"optional" json:"containerSpecGcsPath" yaml:"containerSpecGcsPath"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#environment DataPipelinePipeline#environment}
	Environment *DataPipelinePipelineWorkloadDataflowFlexTemplateRequestLaunchParameterEnvironment `field:"optional" json:"environment" yaml:"environment"`
	// Launch options for this Flex Template job.
	//
	// This is a common set of options across languages and templates. This should not be used to pass job parameters.
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#launch_options DataPipelinePipeline#launch_options}
	LaunchOptions *map[string]*string `field:"optional" json:"launchOptions" yaml:"launchOptions"`
	// 'The parameters for the Flex Template.
	//
	// Example: {"numWorkers":"5"}'
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#parameters DataPipelinePipeline#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// 'Use this to pass transform name mappings for streaming update jobs.
	//
	// Example: {"oldTransformName":"newTransformName",...}'
	// 'An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#transform_name_mappings DataPipelinePipeline#transform_name_mappings}
	TransformNameMappings *map[string]*string `field:"optional" json:"transformNameMappings" yaml:"transformNameMappings"`
	// Set this to true if you are sending a request to update a running streaming job.
	//
	// When set, the job name should be the same as the running job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_pipeline_pipeline#update DataPipelinePipeline#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

