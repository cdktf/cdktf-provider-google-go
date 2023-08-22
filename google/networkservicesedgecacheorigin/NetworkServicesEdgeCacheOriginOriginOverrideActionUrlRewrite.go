package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginOriginOverrideActionUrlRewrite struct {
	// Prior to forwarding the request to the selected origin, the request's host header is replaced with contents of the hostRewrite.
	//
	// This value must be between 1 and 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_services_edge_cache_origin#host_rewrite NetworkServicesEdgeCacheOrigin#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
}

