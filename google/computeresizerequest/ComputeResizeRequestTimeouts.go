// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresizerequest


type ComputeResizeRequestTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_resize_request#create ComputeResizeRequest#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_resize_request#delete ComputeResizeRequest#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

