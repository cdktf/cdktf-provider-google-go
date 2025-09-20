// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package colabschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ColabScheduleConfig struct {
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
	// create_notebook_execution_job_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#create_notebook_execution_job_request ColabSchedule#create_notebook_execution_job_request}
	CreateNotebookExecutionJobRequest *ColabScheduleCreateNotebookExecutionJobRequest `field:"required" json:"createNotebookExecutionJobRequest" yaml:"createNotebookExecutionJobRequest"`
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled runs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#cron ColabSchedule#cron}
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// Required. The display name of the Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#display_name ColabSchedule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource: https://cloud.google.com/colab/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#location ColabSchedule#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Maximum number of runs that can be started concurrently for this Schedule.
	//
	// This is the limit for starting the scheduled requests and not the execution of the notebook execution jobs created by the requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#max_concurrent_run_count ColabSchedule#max_concurrent_run_count}
	MaxConcurrentRunCount *string `field:"required" json:"maxConcurrentRunCount" yaml:"maxConcurrentRunCount"`
	// Whether new scheduled runs can be queued when max_concurrent_runs limit is reached.
	//
	// If set to true, new runs will be queued instead of skipped. Default to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#allow_queueing ColabSchedule#allow_queueing}
	AllowQueueing interface{} `field:"optional" json:"allowQueueing" yaml:"allowQueueing"`
	// Desired state of the Colab Schedule.
	//
	// Set this field to 'ACTIVE' to start/resume the schedule, and 'PAUSED' to pause the schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#desired_state ColabSchedule#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Timestamp after which no new runs can be scheduled.
	//
	// If specified, the schedule will be completed when either end_time is reached or when scheduled_run_count >= max_run_count. Must be in the RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#end_time ColabSchedule#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#id ColabSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maximum run count of the schedule.
	//
	// If specified, The schedule will be completed when either startedRunCount >= maxRunCount or when endTime is reached. If not specified, new runs will keep getting scheduled until this Schedule is paused or deleted. Already scheduled runs will be allowed to complete. Unset if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#max_run_count ColabSchedule#max_run_count}
	MaxRunCount *string `field:"optional" json:"maxRunCount" yaml:"maxRunCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#project ColabSchedule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The timestamp after which the first run can be scheduled.
	//
	// Defaults to the schedule creation time. Must be in the RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#start_time ColabSchedule#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/colab_schedule#timeouts ColabSchedule#timeouts}
	Timeouts *ColabScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

