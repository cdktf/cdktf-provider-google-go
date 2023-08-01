package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteAction struct {
	// cdn_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#cdn_policy NetworkServicesEdgeCacheService#cdn_policy}
	CdnPolicy *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicy `field:"optional" json:"cdnPolicy" yaml:"cdnPolicy"`
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#cors_policy NetworkServicesEdgeCacheService#cors_policy}
	CorsPolicy *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#url_rewrite NetworkServicesEdgeCacheService#url_rewrite}
	UrlRewrite *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
}

