// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetHeaderOverrides struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/cloud_tasks_queue#header CloudTasksQueue#header}
	Header *CloudTasksQueueHttpTargetHeaderOverridesHeader `field:"required" json:"header" yaml:"header"`
}

