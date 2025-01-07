// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetHeaderOverridesHeader struct {
	// The Key of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/cloud_tasks_queue#key CloudTasksQueue#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/cloud_tasks_queue#value CloudTasksQueue#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

