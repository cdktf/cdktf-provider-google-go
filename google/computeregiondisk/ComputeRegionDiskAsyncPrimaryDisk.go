// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiondisk


type ComputeRegionDiskAsyncPrimaryDisk struct {
	// Primary disk for asynchronous disk replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/compute_region_disk#disk ComputeRegionDisk#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

