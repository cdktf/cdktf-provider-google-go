// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computevpngateway


type ComputeVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_vpn_gateway#create ComputeVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_vpn_gateway#delete ComputeVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

