// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionhealthcheck


type ComputeRegionHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/compute_region_health_check#create ComputeRegionHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/compute_region_health_check#delete ComputeRegionHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/compute_region_health_check#update ComputeRegionHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

