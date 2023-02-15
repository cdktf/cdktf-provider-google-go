package computefirewallpolicy


type ComputeFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy#create ComputeFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy#delete ComputeFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_firewall_policy#update ComputeFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

