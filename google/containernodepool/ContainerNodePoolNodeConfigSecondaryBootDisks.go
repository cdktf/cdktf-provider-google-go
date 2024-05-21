// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigSecondaryBootDisks struct {
	// Disk image to create the secondary boot disk from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/container_node_pool#disk_image ContainerNodePool#disk_image}
	DiskImage *string `field:"required" json:"diskImage" yaml:"diskImage"`
	// Mode for how the secondary boot disk is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/container_node_pool#mode ContainerNodePool#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

