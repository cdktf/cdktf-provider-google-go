// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeglobalnetworkendpoint


type ComputeGlobalNetworkEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_global_network_endpoint#create ComputeGlobalNetworkEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_global_network_endpoint#delete ComputeGlobalNetworkEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

