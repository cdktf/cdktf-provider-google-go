package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteAction struct {
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#request_mirror_policy ComputeRegionUrlMap#request_mirror_policy}
	RequestMirrorPolicy *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#retry_policy ComputeRegionUrlMap#retry_policy}
	RetryPolicy *ComputeRegionUrlMapDefaultRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#weighted_backend_services ComputeRegionUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

