// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerStandbyPolicy struct {
	// Specifies the number of seconds that the MIG should wait to suspend or stop a VM after that VM was created.
	//
	// The initial delay gives the initialization script the time to prepare your VM for a quick scale out. The value of initial delay must be between 0 and 3600 seconds. The default value is 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_group_manager#initial_delay_sec ComputeRegionInstanceGroupManager#initial_delay_sec}
	InitialDelaySec *float64 `field:"optional" json:"initialDelaySec" yaml:"initialDelaySec"`
	// Defines how a MIG resumes or starts VMs from a standby pool when the group scales out.
	//
	// The default mode is "MANUAL".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_region_instance_group_manager#mode ComputeRegionInstanceGroupManager#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

