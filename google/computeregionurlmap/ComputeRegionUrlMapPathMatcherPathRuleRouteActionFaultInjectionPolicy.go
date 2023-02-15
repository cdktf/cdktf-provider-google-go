package computeregionurlmap


type ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#abort ComputeRegionUrlMap#abort}
	Abort *ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#delay ComputeRegionUrlMap#delay}
	Delay *ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

