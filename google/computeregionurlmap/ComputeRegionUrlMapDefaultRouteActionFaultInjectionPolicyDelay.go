package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay struct {
	// fixed_delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#fixed_delay ComputeRegionUrlMap#fixed_delay}
	FixedDelay *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelayFixedDelay `field:"optional" json:"fixedDelay" yaml:"fixedDelay"`
	// The percentage of traffic (connections/operations/requests) on which delay will be introduced as part of fault injection.
	//
	// The value must be between 0.0 and 100.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#percentage ComputeRegionUrlMap#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

