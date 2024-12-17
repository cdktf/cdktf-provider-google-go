// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioncommitment


type ComputeRegionCommitmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_region_commitment#create ComputeRegionCommitment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_region_commitment#delete ComputeRegionCommitment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

