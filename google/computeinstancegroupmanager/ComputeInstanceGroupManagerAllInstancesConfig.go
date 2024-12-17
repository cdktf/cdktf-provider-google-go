// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancegroupmanager


type ComputeInstanceGroupManagerAllInstancesConfig struct {
	// The label key-value pairs that you want to patch onto the instance,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_instance_group_manager#labels ComputeInstanceGroupManager#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The metadata key-value pairs that you want to patch onto the instance.
	//
	// For more information, see Project and instance metadata,
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_instance_group_manager#metadata ComputeInstanceGroupManager#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

