package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite struct {
	// Prior to forwarding the request to the selected origin, the request's host header is replaced with contents of hostRewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#host_rewrite NetworkServicesEdgeCacheService#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Prior to forwarding the request to the selected origin, the matching portion of the request's path is replaced by pathPrefixRewrite.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#path_prefix_rewrite NetworkServicesEdgeCacheService#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
	// Prior to forwarding the request to the selected origin, if the request matched a pathTemplateMatch, the matching portion of the request's path is replaced re-written using the pattern specified by pathTemplateRewrite.
	//
	// pathTemplateRewrite must be between 1 and 255 characters
	// (inclusive), must start with a '/', and must only use variables
	// captured by the route's pathTemplate matchers.
	//
	// pathTemplateRewrite may only be used when all of a route's
	// MatchRules specify pathTemplate.
	//
	// Only one of pathPrefixRewrite and pathTemplateRewrite may be
	// specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#path_template_rewrite NetworkServicesEdgeCacheService#path_template_rewrite}
	PathTemplateRewrite *string `field:"optional" json:"pathTemplateRewrite" yaml:"pathTemplateRewrite"`
}

