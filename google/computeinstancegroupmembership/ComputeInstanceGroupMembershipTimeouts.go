// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupmembership


type ComputeInstanceGroupMembershipTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_group_membership#create ComputeInstanceGroupMembership#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/compute_instance_group_membership#delete ComputeInstanceGroupMembership#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

