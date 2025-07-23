// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigSecondaryBootDisks struct {
	// Disk image to create the secondary boot disk from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#disk_image ContainerCluster#disk_image}
	DiskImage *string `field:"required" json:"diskImage" yaml:"diskImage"`
	// Mode for how the secondary boot disk is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#mode ContainerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

