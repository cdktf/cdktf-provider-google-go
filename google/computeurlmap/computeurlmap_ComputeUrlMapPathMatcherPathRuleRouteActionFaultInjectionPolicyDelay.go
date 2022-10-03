package computeurlmap


type ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay struct {
	// fixed_delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_url_map#fixed_delay ComputeUrlMap#fixed_delay}
	FixedDelay *ComputeUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelayFixedDelay `field:"required" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.
	//
	// The value must be between 0.0 and
	// 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_url_map#percentage ComputeUrlMap#percentage}
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

