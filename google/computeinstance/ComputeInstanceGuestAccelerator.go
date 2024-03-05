// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstance


type ComputeInstanceGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/compute_instance#count ComputeInstance#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/compute_instance#type ComputeInstance#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

