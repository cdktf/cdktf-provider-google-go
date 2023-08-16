package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyAddSignatures struct {
	// The actions to take to add signatures to responses. Possible values: ["GENERATE_COOKIE", "GENERATE_TOKEN_HLS_COOKIELESS", "PROPAGATE_TOKEN_HLS_COOKIELESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#actions NetworkServicesEdgeCacheService#actions}
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// The parameters to copy from the verified token to the generated token.
	//
	// Only the following parameters may be copied:
	//
	// 'PathGlobs'
	// 'paths'
	// 'acl'
	// 'URLPrefix'
	// 'IPRanges'
	// 'SessionID'
	// 'id'
	// 'Data'
	// 'data'
	// 'payload'
	// 'Headers'
	//
	// You may specify up to 6 parameters to copy.  A given parameter is be copied only if the parameter exists in the verified token.  Parameter names are matched exactly as specified.  The order of the parameters does not matter.  Duplicates are not allowed.
	//
	// This field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#copied_parameters NetworkServicesEdgeCacheService#copied_parameters}
	CopiedParameters *[]*string `field:"optional" json:"copiedParameters" yaml:"copiedParameters"`
	// The keyset to use for signature generation.
	//
	// The following are both valid paths to an EdgeCacheKeyset resource:
	//
	// 'projects/project/locations/global/edgeCacheKeysets/yourKeyset'
	// 'yourKeyset'
	//
	// This must be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.  This field may not be specified otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#keyset NetworkServicesEdgeCacheService#keyset}
	Keyset *string `field:"optional" json:"keyset" yaml:"keyset"`
	// The query parameter in which to put the generated token.
	//
	// If not specified, defaults to 'edge-cache-token'.
	//
	// If specified, the name must be 1-64 characters long and match the regular expression '[a-zA-Z]([a-zA-Z0-9_-])*' which means the first character must be a letter, and all following characters must be a dash, underscore, letter or digit.
	//
	// This field may only be set when the GENERATE_TOKEN_HLS_COOKIELESS or PROPAGATE_TOKEN_HLS_COOKIELESS actions are specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#token_query_parameter NetworkServicesEdgeCacheService#token_query_parameter}
	TokenQueryParameter *string `field:"optional" json:"tokenQueryParameter" yaml:"tokenQueryParameter"`
	// The duration the token is valid starting from the moment the token is first generated.
	//
	// Defaults to '86400s' (1 day).
	//
	// The TTL must be >= 0 and <= 604,800 seconds (1 week).
	//
	// This field may only be specified when the GENERATE_COOKIE or GENERATE_TOKEN_HLS_COOKIELESS actions are specified.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#token_ttl NetworkServicesEdgeCacheService#token_ttl}
	TokenTtl *string `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
}

