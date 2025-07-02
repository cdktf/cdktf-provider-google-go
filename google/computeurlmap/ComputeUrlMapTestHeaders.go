// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapTestHeaders struct {
	// Header name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_url_map#name ComputeUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Header value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/compute_url_map#value ComputeUrlMap#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

