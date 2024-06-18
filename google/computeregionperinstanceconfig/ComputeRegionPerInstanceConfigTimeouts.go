// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionperinstanceconfig


type ComputeRegionPerInstanceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/compute_region_per_instance_config#create ComputeRegionPerInstanceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/compute_region_per_instance_config#delete ComputeRegionPerInstanceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.34.0/docs/resources/compute_region_per_instance_config#update ComputeRegionPerInstanceConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

