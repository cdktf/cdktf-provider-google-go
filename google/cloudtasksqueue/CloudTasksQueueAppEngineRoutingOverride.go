// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueAppEngineRoutingOverride struct {
	// App instance.
	//
	// By default, the task is sent to an instance which is available when the task is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/cloud_tasks_queue#instance CloudTasksQueue#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// App service.
	//
	// By default, the task is sent to the service which is the default service when the task is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/cloud_tasks_queue#service CloudTasksQueue#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// App version.
	//
	// By default, the task is sent to the version which is the default version when the task is attempted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/cloud_tasks_queue#version CloudTasksQueue#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

