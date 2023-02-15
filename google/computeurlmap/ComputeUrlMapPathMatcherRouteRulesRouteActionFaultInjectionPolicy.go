package computeurlmap


type ComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_url_map#abort ComputeUrlMap#abort}
	Abort *ComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_url_map#delay ComputeUrlMap#delay}
	Delay *ComputeUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

