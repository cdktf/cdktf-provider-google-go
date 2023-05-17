package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAdd struct {
	// The name of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_services_edge_cache_service#header_name NetworkServicesEdgeCacheService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_services_edge_cache_service#header_value NetworkServicesEdgeCacheService#header_value}
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
	// Whether to replace all existing headers with the same name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/network_services_edge_cache_service#replace NetworkServicesEdgeCacheService#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
}

