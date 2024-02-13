// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehavpngateway


type ComputeHaVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_ha_vpn_gateway#create ComputeHaVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_ha_vpn_gateway#delete ComputeHaVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

