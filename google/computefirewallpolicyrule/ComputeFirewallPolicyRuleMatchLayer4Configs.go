package computefirewallpolicyrule


type ComputeFirewallPolicyRuleMatchLayer4Configs struct {
	// The IP protocol to which this rule applies.
	//
	// The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (`tcp`, `udp`, `icmp`, `esp`, `ah`, `ipip`, `sctp`), or the IP protocol number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_firewall_policy_rule#ip_protocol ComputeFirewallPolicyRule#ip_protocol}
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// An optional list of ports to which this rule applies.
	//
	// This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_firewall_policy_rule#ports ComputeFirewallPolicyRule#ports}
	Ports *[]*string `field:"optional" json:"ports" yaml:"ports"`
}

