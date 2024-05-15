// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DataPipelinePipelineScheduleInfo struct {
	// Unix-cron format of the schedule. This information is retrieved from the linked Cloud Scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/data_pipeline_pipeline#schedule DataPipelinePipeline#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// Timezone ID. This matches the timezone IDs used by the Cloud Scheduler API. If empty, UTC time is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/data_pipeline_pipeline#time_zone DataPipelinePipeline#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

