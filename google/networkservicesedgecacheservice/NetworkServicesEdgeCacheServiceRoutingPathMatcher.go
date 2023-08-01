package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcher struct {
	// The name to which this PathMatcher is referred by the HostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#name NetworkServicesEdgeCacheService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// route_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#route_rule NetworkServicesEdgeCacheService#route_rule}
	RouteRule interface{} `field:"required" json:"routeRule" yaml:"routeRule"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/network_services_edge_cache_service#description NetworkServicesEdgeCacheService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

