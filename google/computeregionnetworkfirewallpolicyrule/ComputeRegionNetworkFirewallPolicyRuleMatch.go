// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicyrule


type ComputeRegionNetworkFirewallPolicyRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#layer4_configs ComputeRegionNetworkFirewallPolicyRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#dest_address_groups ComputeRegionNetworkFirewallPolicyRule#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic destination.
	//
	// Maximum number of destination fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#dest_fqdns ComputeRegionNetworkFirewallPolicyRule#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#dest_ip_ranges ComputeRegionNetworkFirewallPolicyRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// Region codes whose IP addresses will be used to match for destination of traffic.
	//
	// Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of dest region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#dest_region_codes ComputeRegionNetworkFirewallPolicyRule#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#dest_threat_intelligences ComputeRegionNetworkFirewallPolicyRule#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_address_groups ComputeRegionNetworkFirewallPolicyRule#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Fully Qualified Domain Name (FQDN) which should be matched against traffic source.
	//
	// Maximum number of source fqdn allowed is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_fqdns ComputeRegionNetworkFirewallPolicyRule#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_ip_ranges ComputeRegionNetworkFirewallPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// Region codes whose IP addresses will be used to match for source of traffic.
	//
	// Should be specified as 2 letter country code defined as per ISO 3166 alpha-2 country codes. ex."US" Maximum number of source region codes allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_region_codes ComputeRegionNetworkFirewallPolicyRule#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// src_secure_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_secure_tags ComputeRegionNetworkFirewallPolicyRule#src_secure_tags}
	SrcSecureTags interface{} `field:"optional" json:"srcSecureTags" yaml:"srcSecureTags"`
	// Names of Network Threat Intelligence lists. The IPs in these lists will be matched against traffic source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_network_firewall_policy_rule#src_threat_intelligences ComputeRegionNetworkFirewallPolicyRule#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

