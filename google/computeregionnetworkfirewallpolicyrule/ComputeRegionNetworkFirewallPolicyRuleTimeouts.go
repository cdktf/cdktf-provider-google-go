// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicyrule


type ComputeRegionNetworkFirewallPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/compute_region_network_firewall_policy_rule#create ComputeRegionNetworkFirewallPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/compute_region_network_firewall_policy_rule#delete ComputeRegionNetworkFirewallPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/compute_region_network_firewall_policy_rule#update ComputeRegionNetworkFirewallPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

