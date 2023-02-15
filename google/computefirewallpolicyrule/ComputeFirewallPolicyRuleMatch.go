package computefirewallpolicyrule


type ComputeFirewallPolicyRuleMatch struct {
	// layer4_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy_rule#layer4_configs ComputeFirewallPolicyRule#layer4_configs}
	Layer4Configs interface{} `field:"required" json:"layer4Configs" yaml:"layer4Configs"`
	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy_rule#dest_ip_ranges ComputeFirewallPolicyRule#dest_ip_ranges}
	DestIpRanges *[]*string `field:"optional" json:"destIpRanges" yaml:"destIpRanges"`
	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy_rule#src_ip_ranges ComputeFirewallPolicyRule#src_ip_ranges}
	SrcIpRanges *[]*string `field:"optional" json:"srcIpRanges" yaml:"srcIpRanges"`
}

