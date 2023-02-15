package computeforwardingrule


type ComputeForwardingRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#create ComputeForwardingRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#delete ComputeForwardingRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_forwarding_rule#update ComputeForwardingRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

