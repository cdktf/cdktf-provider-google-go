// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupmanager


type ComputeInstanceGroupManagerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/compute_instance_group_manager#create ComputeInstanceGroupManager#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/compute_instance_group_manager#delete ComputeInstanceGroupManager#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/compute_instance_group_manager#update ComputeInstanceGroupManager#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

