// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkendpointgroup


type ComputeNetworkEndpointGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_network_endpoint_group#create ComputeNetworkEndpointGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_network_endpoint_group#delete ComputeNetworkEndpointGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

