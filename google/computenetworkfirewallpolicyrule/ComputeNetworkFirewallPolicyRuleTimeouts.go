package computenetworkfirewallpolicyrule


type ComputeNetworkFirewallPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_network_firewall_policy_rule#create ComputeNetworkFirewallPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_network_firewall_policy_rule#delete ComputeNetworkFirewallPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_network_firewall_policy_rule#update ComputeNetworkFirewallPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

