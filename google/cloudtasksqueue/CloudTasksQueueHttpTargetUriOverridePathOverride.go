// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetUriOverridePathOverride struct {
	// The URI path (e.g., /users/1234). Default is an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/cloud_tasks_queue#path CloudTasksQueue#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

