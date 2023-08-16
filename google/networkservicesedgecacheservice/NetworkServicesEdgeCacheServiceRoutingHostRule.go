package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingHostRule struct {
	// The list of host patterns to match.
	//
	// Host patterns must be valid hostnames. Ports are not allowed. Wildcard hosts are supported in the suffix or prefix form. * matches any string of ([a-z0-9-.]*). It does not match the empty string.
	//
	// When multiple hosts are specified, hosts are matched in the following priority:
	//
	// 1. Exact domain names: ''www.foo.com''.
	// 2. Suffix domain wildcards: ''*.foo.com'' or ''*-bar.foo.com''.
	// 3. Prefix domain wildcards: ''foo.*'' or ''foo-*''.
	// 4. Special wildcard ''*'' matching any domain.
	//
	// Notes:
	//
	// The wildcard will not match the empty string. e.g. ''*-bar.foo.com'' will match ''baz-bar.foo.com'' but not ''-bar.foo.com''. The longest wildcards match first. Only a single host in the entire service can match on ''*''. A domain must be unique across all configured hosts within a service.
	//
	// Hosts are matched against the HTTP Host header, or for HTTP/2 and HTTP/3, the ":authority" header, from the incoming request.
	//
	// You may specify up to 10 hosts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#hosts NetworkServicesEdgeCacheService#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// The name of the pathMatcher associated with this hostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#path_matcher NetworkServicesEdgeCacheService#path_matcher}
	PathMatcher *string `field:"required" json:"pathMatcher" yaml:"pathMatcher"`
	// A human-readable description of the hostRule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#description NetworkServicesEdgeCacheService#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

