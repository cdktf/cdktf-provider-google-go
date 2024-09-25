// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstance


type ComputeInstanceSchedulingOnInstanceStopAction struct {
	// If true, the contents of any attached Local SSD disks will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/compute_instance#discard_local_ssd ComputeInstance#discard_local_ssd}
	DiscardLocalSsd interface{} `field:"optional" json:"discardLocalSsd" yaml:"discardLocalSsd"`
}

