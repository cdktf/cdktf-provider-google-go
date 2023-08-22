package computeregionurlmap


type ComputeRegionUrlMapPathMatcherPathRuleRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#cors_policy ComputeRegionUrlMap#cors_policy}
	CorsPolicy *ComputeRegionUrlMapPathMatcherPathRuleRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#fault_injection_policy ComputeRegionUrlMap#fault_injection_policy}
	FaultInjectionPolicy *ComputeRegionUrlMapPathMatcherPathRuleRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#request_mirror_policy ComputeRegionUrlMap#request_mirror_policy}
	RequestMirrorPolicy *ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#retry_policy ComputeRegionUrlMap#retry_policy}
	RetryPolicy *ComputeRegionUrlMapPathMatcherPathRuleRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#timeout ComputeRegionUrlMap#timeout}
	Timeout *ComputeRegionUrlMapPathMatcherPathRuleRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#url_rewrite ComputeRegionUrlMap#url_rewrite}
	UrlRewrite *ComputeRegionUrlMapPathMatcherPathRuleRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#weighted_backend_services ComputeRegionUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

