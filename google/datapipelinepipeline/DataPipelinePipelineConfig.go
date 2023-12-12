// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataPipelinePipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// "The pipeline name.
	//
	// For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
	// "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
	// "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
	// "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#name DataPipelinePipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The state of the pipeline.
	//
	// When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
	// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state Possible values: ["STATE_UNSPECIFIED", "STATE_RESUMING", "STATE_ACTIVE", "STATE_STOPPING", "STATE_ARCHIVED", "STATE_PAUSED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#state DataPipelinePipeline#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The type of the pipeline.
	//
	// This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
	// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype Possible values: ["PIPELINE_TYPE_UNSPECIFIED", "PIPELINE_TYPE_BATCH", "PIPELINE_TYPE_STREAMING"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#type DataPipelinePipeline#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#display_name DataPipelinePipeline#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#id DataPipelinePipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The sources of the pipeline (for example, Dataplex).
	//
	// The keys and values are set by the corresponding sources during pipeline creation.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#pipeline_sources DataPipelinePipeline#pipeline_sources}
	PipelineSources *map[string]*string `field:"optional" json:"pipelineSources" yaml:"pipelineSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#project DataPipelinePipeline#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#region DataPipelinePipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// schedule_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#schedule_info DataPipelinePipeline#schedule_info}
	ScheduleInfo *DataPipelinePipelineScheduleInfo `field:"optional" json:"scheduleInfo" yaml:"scheduleInfo"`
	// Optional.
	//
	// A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#scheduler_service_account_email DataPipelinePipeline#scheduler_service_account_email}
	SchedulerServiceAccountEmail *string `field:"optional" json:"schedulerServiceAccountEmail" yaml:"schedulerServiceAccountEmail"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#timeouts DataPipelinePipeline#timeouts}
	Timeouts *DataPipelinePipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// workload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/data_pipeline_pipeline#workload DataPipelinePipeline#workload}
	Workload *DataPipelinePipelineWorkload `field:"optional" json:"workload" yaml:"workload"`
}

