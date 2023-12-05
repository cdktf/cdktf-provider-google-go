// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computediskasyncreplication


type ComputeDiskAsyncReplicationSecondaryDisk struct {
	// Secondary disk for asynchronous replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.8.0/docs/resources/compute_disk_async_replication#disk ComputeDiskAsyncReplication#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
}

