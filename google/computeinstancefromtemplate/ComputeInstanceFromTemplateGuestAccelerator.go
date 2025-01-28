// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate


type ComputeInstanceFromTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_instance_from_template#count ComputeInstanceFromTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource exposed to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_instance_from_template#type ComputeInstanceFromTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

