// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupnamedport


type ComputeInstanceGroupNamedPortTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/compute_instance_group_named_port#create ComputeInstanceGroupNamedPortA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/compute_instance_group_named_port#delete ComputeInstanceGroupNamedPortA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

