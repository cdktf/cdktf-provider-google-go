// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerStatefulExternalIp struct {
	// A value that prescribes what should happen to an associated static Address resource when a VM instance is permanently deleted.
	//
	// The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION. NEVER - detach the IP when the VM is deleted, but do not delete the address resource. ON_PERMANENT_INSTANCE_DELETION will delete the stateful address when the VM is permanently deleted from the instance group. The default is NEVER.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_region_instance_group_manager#delete_rule ComputeRegionInstanceGroupManager#delete_rule}
	DeleteRule *string `field:"optional" json:"deleteRule" yaml:"deleteRule"`
	// The network interface name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_region_instance_group_manager#interface_name ComputeRegionInstanceGroupManager#interface_name}
	InterfaceName *string `field:"optional" json:"interfaceName" yaml:"interfaceName"`
}

