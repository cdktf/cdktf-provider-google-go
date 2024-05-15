// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute


type NetworkConnectivityPolicyBasedRouteFilter struct {
	// Internet protocol versions this policy-based route applies to. Possible values: ["IPV4"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/network_connectivity_policy_based_route#protocol_version NetworkConnectivityPolicyBasedRoute#protocol_version}
	ProtocolVersion *string `field:"required" json:"protocolVersion" yaml:"protocolVersion"`
	// The destination IP range of outgoing packets that this policy-based route applies to.
	//
	// Default is "0.0.0.0/0" if protocol version is IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/network_connectivity_policy_based_route#dest_range NetworkConnectivityPolicyBasedRoute#dest_range}
	DestRange *string `field:"optional" json:"destRange" yaml:"destRange"`
	// The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/network_connectivity_policy_based_route#ip_protocol NetworkConnectivityPolicyBasedRoute#ip_protocol}
	IpProtocol *string `field:"optional" json:"ipProtocol" yaml:"ipProtocol"`
	// The source IP range of outgoing packets that this policy-based route applies to.
	//
	// Default is "0.0.0.0/0" if protocol version is IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.1/docs/resources/network_connectivity_policy_based_route#src_range NetworkConnectivityPolicyBasedRoute#src_range}
	SrcRange *string `field:"optional" json:"srcRange" yaml:"srcRange"`
}

