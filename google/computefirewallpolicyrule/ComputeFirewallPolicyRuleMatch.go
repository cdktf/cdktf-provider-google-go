package computefirewallpolicyrule


type ComputeFirewallPolicyRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#layer4_configs ComputeFirewallPolicyRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// Address groups which should be matched against the traffic destination.
	//
	// Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#dest_address_groups ComputeFirewallPolicyRule#dest_address_groups}
	DestAddressGroups *[]*string `field:"optional" json:"destAddressGroups" yaml:"destAddressGroups"`
	// Domain names that will be used to match against the resolved domain name of destination of traffic.
	//
	// Can only be specified if DIRECTION is egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#dest_fqdns ComputeFirewallPolicyRule#dest_fqdns}
	DestFqdns *[]*string `field:"optional" json:"destFqdns" yaml:"destFqdns"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#dest_ip_ranges ComputeFirewallPolicyRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// The Unicode country codes whose IP addresses will be used to match against the source of traffic.
	//
	// Can only be specified if DIRECTION is egress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#dest_region_codes ComputeFirewallPolicyRule#dest_region_codes}
	DestRegionCodes *[]*string `field:"optional" json:"destRegionCodes" yaml:"destRegionCodes"`
	// Name of the Google Cloud Threat Intelligence list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#dest_threat_intelligences ComputeFirewallPolicyRule#dest_threat_intelligences}
	DestThreatIntelligences *[]*string `field:"optional" json:"destThreatIntelligences" yaml:"destThreatIntelligences"`
	// Address groups which should be matched against the traffic source.
	//
	// Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#src_address_groups ComputeFirewallPolicyRule#src_address_groups}
	SrcAddressGroups *[]*string `field:"optional" json:"srcAddressGroups" yaml:"srcAddressGroups"`
	// Domain names that will be used to match against the resolved domain name of source of traffic.
	//
	// Can only be specified if DIRECTION is ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#src_fqdns ComputeFirewallPolicyRule#src_fqdns}
	SrcFqdns *[]*string `field:"optional" json:"srcFqdns" yaml:"srcFqdns"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#src_ip_ranges ComputeFirewallPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
	// The Unicode country codes whose IP addresses will be used to match against the source of traffic.
	//
	// Can only be specified if DIRECTION is ingress.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#src_region_codes ComputeFirewallPolicyRule#src_region_codes}
	SrcRegionCodes *[]*string `field:"optional" json:"srcRegionCodes" yaml:"srcRegionCodes"`
	// Name of the Google Cloud Threat Intelligence list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_firewall_policy_rule#src_threat_intelligences ComputeFirewallPolicyRule#src_threat_intelligences}
	SrcThreatIntelligences *[]*string `field:"optional" json:"srcThreatIntelligences" yaml:"srcThreatIntelligences"`
}

