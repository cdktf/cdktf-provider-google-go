// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate


type ComputeRegionInstanceTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_region_instance_template#count ComputeRegionInstanceTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_region_instance_template#type ComputeRegionInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

