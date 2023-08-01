package computeforwardingrule


type ComputeForwardingRuleServiceDirectoryRegistrations struct {
	// Service Directory namespace to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_forwarding_rule#namespace ComputeForwardingRule#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Service Directory service to register the forwarding rule under.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_forwarding_rule#service ComputeForwardingRule#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

