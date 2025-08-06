// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityfirewallendpoint


type NetworkSecurityFirewallEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_security_firewall_endpoint#create NetworkSecurityFirewallEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_security_firewall_endpoint#delete NetworkSecurityFirewallEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/network_security_firewall_endpoint#update NetworkSecurityFirewallEndpoint#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

