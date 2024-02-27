// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computediskasyncreplication


type ComputeDiskAsyncReplicationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_disk_async_replication#create ComputeDiskAsyncReplication#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_disk_async_replication#delete ComputeDiskAsyncReplication#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

