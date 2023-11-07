// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicy


type ComputeRegionNetworkFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/compute_region_network_firewall_policy#create ComputeRegionNetworkFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/compute_region_network_firewall_policy#delete ComputeRegionNetworkFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/compute_region_network_firewall_policy#update ComputeRegionNetworkFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

