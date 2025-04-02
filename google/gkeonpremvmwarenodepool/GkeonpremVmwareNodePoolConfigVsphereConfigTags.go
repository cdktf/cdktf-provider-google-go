// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarenodepool


type GkeonpremVmwareNodePoolConfigVsphereConfigTags struct {
	// The Vsphere tag category.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gkeonprem_vmware_node_pool#category GkeonpremVmwareNodePool#category}
	Category *string `field:"optional" json:"category" yaml:"category"`
	// The Vsphere tag name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gkeonprem_vmware_node_pool#tag GkeonpremVmwareNodePool#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

