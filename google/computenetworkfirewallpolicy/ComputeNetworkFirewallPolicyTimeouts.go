// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicy


type ComputeNetworkFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_network_firewall_policy#create ComputeNetworkFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_network_firewall_policy#delete ComputeNetworkFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/compute_network_firewall_policy#update ComputeNetworkFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

