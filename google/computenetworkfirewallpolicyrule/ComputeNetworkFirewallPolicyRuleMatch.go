package computenetworkfirewallpolicyrule


type ComputeNetworkFirewallPolicyRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#layer4_configs ComputeNetworkFirewallPolicyRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// Address groups which should be matched against the traffic destination.
	//
	// Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#dest_address_groups ComputeNetworkFirewallPolicyRule#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Domain names that will be used to match against the resolved domain name of destination of traffic.
	//
	// Can only be specified if DIRECTION is egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#dest_fqdns ComputeNetworkFirewallPolicyRule#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#dest_ip_ranges ComputeNetworkFirewallPolicyRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// The Unicode country codes whose IP addresses will be used to match against the source of traffic.
	//
	// Can only be specified if DIRECTION is egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#dest_region_codes ComputeNetworkFirewallPolicyRule#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Name of the Google Cloud Threat Intelligence list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#dest_threat_intelligences ComputeNetworkFirewallPolicyRule#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source.
	//
	// Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_address_groups ComputeNetworkFirewallPolicyRule#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Domain names that will be used to match against the resolved domain name of source of traffic.
	//
	// Can only be specified if DIRECTION is ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_fqdns ComputeNetworkFirewallPolicyRule#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 5000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_ip_ranges ComputeNetworkFirewallPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// The Unicode country codes whose IP addresses will be used to match against the source of traffic.
	//
	// Can only be specified if DIRECTION is ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_region_codes ComputeNetworkFirewallPolicyRule#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// src_secure_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_secure_tags ComputeNetworkFirewallPolicyRule#src_secure_tags}
	SrcSecureTags interface{} `field:"optional" json:"srcSecureTags" yaml:"srcSecureTags"`
	// Name of the Google Cloud Threat Intelligence list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_firewall_policy_rule#src_threat_intelligences ComputeNetworkFirewallPolicyRule#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

