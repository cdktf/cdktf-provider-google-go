// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeexternalvpngateway


type ComputeExternalVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/compute_external_vpn_gateway#create ComputeExternalVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/compute_external_vpn_gateway#delete ComputeExternalVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/compute_external_vpn_gateway#update ComputeExternalVpnGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

