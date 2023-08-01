package computeregionurlmap


type ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#abort ComputeRegionUrlMap#abort}
	Abort *ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#delay ComputeRegionUrlMap#delay}
	Delay *ComputeRegionUrlMapPathMatcherRouteRulesRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

