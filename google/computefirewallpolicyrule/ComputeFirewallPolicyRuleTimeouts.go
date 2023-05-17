package computefirewallpolicyrule


type ComputeFirewallPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_firewall_policy_rule#create ComputeFirewallPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_firewall_policy_rule#delete ComputeFirewallPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_firewall_policy_rule#update ComputeFirewallPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

