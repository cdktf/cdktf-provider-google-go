// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy struct {
	// instance_selections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_instance_group_manager#instance_selections ComputeRegionInstanceGroupManager#instance_selections}
	InstanceSelections interface{} `field:"optional" json:"instanceSelections" yaml:"instanceSelections"`
}

