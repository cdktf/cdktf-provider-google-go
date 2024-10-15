// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancetemplate


type ComputeInstanceTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/compute_instance_template#count ComputeInstanceTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/compute_instance_template#type ComputeInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

