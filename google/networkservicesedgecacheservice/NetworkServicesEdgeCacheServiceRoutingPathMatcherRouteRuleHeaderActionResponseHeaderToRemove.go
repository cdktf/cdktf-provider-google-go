package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemove struct {
	// Headers to remove from the response prior to sending it back to the client.
	//
	// Response headers are only sent to the client, and do not have an effect on the cache serving the response.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_service#header_name NetworkServicesEdgeCacheService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

