// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicyassociation


type ComputeNetworkFirewallPolicyAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_network_firewall_policy_association#create ComputeNetworkFirewallPolicyAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_network_firewall_policy_association#delete ComputeNetworkFirewallPolicyAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

