// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkendpoints


type ComputeNetworkEndpointsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/compute_network_endpoints#create ComputeNetworkEndpoints#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/compute_network_endpoints#delete ComputeNetworkEndpoints#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/compute_network_endpoints#update ComputeNetworkEndpoints#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

