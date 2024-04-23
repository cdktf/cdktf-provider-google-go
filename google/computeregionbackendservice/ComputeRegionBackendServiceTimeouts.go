// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/compute_region_backend_service#create ComputeRegionBackendService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/compute_region_backend_service#delete ComputeRegionBackendService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.26.0/docs/resources/compute_region_backend_service#update ComputeRegionBackendService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

