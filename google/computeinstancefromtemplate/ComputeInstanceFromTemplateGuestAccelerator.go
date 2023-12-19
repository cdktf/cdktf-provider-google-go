// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/compute_instance_from_template#count ComputeInstanceFromTemplate#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/compute_instance_from_template#type ComputeInstanceFromTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

