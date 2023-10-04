// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionperinstanceconfig


type ComputeRegionPerInstanceConfigPreservedState struct {
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.0.0/docs/resources/compute_region_per_instance_config#disk ComputeRegionPerInstanceConfig#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Preserved metadata defined for this instance. This is a list of key->value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.0.0/docs/resources/compute_region_per_instance_config#metadata ComputeRegionPerInstanceConfig#metadata}
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
}

