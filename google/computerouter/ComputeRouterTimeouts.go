// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouter


type ComputeRouterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router#create ComputeRouter#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router#delete ComputeRouter#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/compute_router#update ComputeRouter#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

