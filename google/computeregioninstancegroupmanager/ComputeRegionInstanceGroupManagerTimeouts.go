// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/compute_region_instance_group_manager#create ComputeRegionInstanceGroupManager#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/compute_region_instance_group_manager#delete ComputeRegionInstanceGroupManager#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/compute_region_instance_group_manager#update ComputeRegionInstanceGroupManager#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

