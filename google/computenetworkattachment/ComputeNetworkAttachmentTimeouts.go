// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkattachment


type ComputeNetworkAttachmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_network_attachment#create ComputeNetworkAttachment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_network_attachment#delete ComputeNetworkAttachment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_network_attachment#update ComputeNetworkAttachment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

