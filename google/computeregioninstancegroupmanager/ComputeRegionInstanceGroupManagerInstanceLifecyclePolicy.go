// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy struct {
	// Specifies whether to apply the group's latest configuration when repairing a VM.
	//
	// Valid options are: YES, NO. If YES and you updated the group's instance template or per-instance configurations after the VM was created, then these changes are applied when VM is repaired. If NO (default), then updates are applied in accordance with the group's update policy type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_region_instance_group_manager#force_update_on_repair ComputeRegionInstanceGroupManager#force_update_on_repair}
	ForceUpdateOnRepair *string `field:"optional" json:"forceUpdateOnRepair" yaml:"forceUpdateOnRepair"`
}

