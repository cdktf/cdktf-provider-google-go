// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computepublicdelegatedprefix


type ComputePublicDelegatedPrefixTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_public_delegated_prefix#create ComputePublicDelegatedPrefix#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_public_delegated_prefix#delete ComputePublicDelegatedPrefix#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

