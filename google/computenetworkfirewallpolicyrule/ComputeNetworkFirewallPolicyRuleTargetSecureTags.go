// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicyrule


type ComputeNetworkFirewallPolicyRuleTargetSecureTags struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_network_firewall_policy_rule#name ComputeNetworkFirewallPolicyRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

