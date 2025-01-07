// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyInstanceSelections struct {
	// Full machine-type names, e.g. "n1-standard-16".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_region_instance_group_manager#machine_types ComputeRegionInstanceGroupManager#machine_types}
	MachineTypes *[]*string `field:"required" json:"machineTypes" yaml:"machineTypes"`
	// Instance selection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_region_instance_group_manager#name ComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Preference of this instance selection.
	//
	// Lower number means higher preference. MIG will first try to create a VM based on the machine-type with lowest rank and fallback to next rank based on availability. Machine types and instance selections with the same rank have the same preference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_region_instance_group_manager#rank ComputeRegionInstanceGroupManager#rank}
	Rank *float64 `field:"optional" json:"rank" yaml:"rank"`
}

