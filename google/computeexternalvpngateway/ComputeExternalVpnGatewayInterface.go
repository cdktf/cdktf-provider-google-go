// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeexternalvpngateway


type ComputeExternalVpnGatewayInterface struct {
	// The numeric ID for this interface.
	//
	// Allowed values are based on the redundancy type
	// of this external VPN gateway
	// * '0 - SINGLE_IP_INTERNALLY_REDUNDANT'
	// * '0, 1 - TWO_IPS_REDUNDANCY'
	// * '0, 1, 2, 3 - FOUR_IPS_REDUNDANCY'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/compute_external_vpn_gateway#id ComputeExternalVpnGateway#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// IP address of the interface in the external VPN gateway.
	//
	// Only IPv4 is supported. This IP address can be either from
	// your on-premise gateway or another Cloud provider's VPN gateway,
	// it cannot be an IP address from Google Compute Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/compute_external_vpn_gateway#ip_address ComputeExternalVpnGateway#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// IPv6 address of the interface in the external VPN gateway.
	//
	// This IPv6
	// address can be either from your on-premise gateway or another Cloud
	// provider's VPN gateway, it cannot be an IP address from Google Compute
	// Engine. Must specify an IPv6 address (not IPV4-mapped) using any format
	// described in RFC 4291 (e.g. 2001:db8:0:0:2d9:51:0:0). The output format
	// is RFC 5952 format (e.g. 2001:db8::2d9:51:0:0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/compute_external_vpn_gateway#ipv6_address ComputeExternalVpnGateway#ipv6_address}
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}

