// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkendpoint


type ComputeRegionNetworkEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_region_network_endpoint#create ComputeRegionNetworkEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_region_network_endpoint#delete ComputeRegionNetworkEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

