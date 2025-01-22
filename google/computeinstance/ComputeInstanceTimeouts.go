// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstance


type ComputeInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_instance#create ComputeInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_instance#delete ComputeInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_instance#update ComputeInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

