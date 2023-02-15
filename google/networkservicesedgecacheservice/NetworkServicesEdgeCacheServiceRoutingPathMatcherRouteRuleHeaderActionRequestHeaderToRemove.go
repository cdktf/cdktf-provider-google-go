package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemove struct {
	// The name of the header to remove.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_service#header_name NetworkServicesEdgeCacheService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

