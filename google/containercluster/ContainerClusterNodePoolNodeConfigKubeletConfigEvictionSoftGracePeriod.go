// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigKubeletConfigEvictionSoftGracePeriod struct {
	// Defines grace period for the imagefs.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#imagefs_available ContainerCluster#imagefs_available}
	ImagefsAvailable *string `field:"optional" json:"imagefsAvailable" yaml:"imagefsAvailable"`
	// Defines grace period for the imagefs.inodesFree soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#imagefs_inodes_free ContainerCluster#imagefs_inodes_free}
	ImagefsInodesFree *string `field:"optional" json:"imagefsInodesFree" yaml:"imagefsInodesFree"`
	// Defines grace period for the memory.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#memory_available ContainerCluster#memory_available}
	MemoryAvailable *string `field:"optional" json:"memoryAvailable" yaml:"memoryAvailable"`
	// Defines grace period for the nodefs.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#nodefs_available ContainerCluster#nodefs_available}
	NodefsAvailable *string `field:"optional" json:"nodefsAvailable" yaml:"nodefsAvailable"`
	// Defines grace period for the nodefs.inodesFree soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#nodefs_inodes_free ContainerCluster#nodefs_inodes_free}
	NodefsInodesFree *string `field:"optional" json:"nodefsInodesFree" yaml:"nodefsInodesFree"`
	// Defines grace period for the pid.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#pid_available ContainerCluster#pid_available}
	PidAvailable *string `field:"optional" json:"pidAvailable" yaml:"pidAvailable"`
}

