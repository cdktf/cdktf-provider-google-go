package networkservicesedgecacheservice


type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCorsPolicy struct {
	// Specifies how long results of a preflight request can be cached by a client in seconds.
	//
	// Note that many browser clients enforce a maximum TTL of 600s (10 minutes).
	//
	// - Setting the value to -1 forces a pre-flight check for all requests (not recommended)
	// - A maximum TTL of 86400s can be set, but note that (as above) some clients may force pre-flight checks at a more regular interval.
	// - This translates to the Access-Control-Max-Age header.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#max_age NetworkServicesEdgeCacheService#max_age}
	MaxAge *string `field:"required" json:"maxAge" yaml:"maxAge"`
	// In response to a preflight request, setting this to true indicates that the actual request can include user credentials.
	//
	// This translates to the Access-Control-Allow-Credentials response header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#allow_credentials NetworkServicesEdgeCacheService#allow_credentials}
	AllowCredentials interface{} `field:"optional" json:"allowCredentials" yaml:"allowCredentials"`
	// Specifies the content for the Access-Control-Allow-Headers response header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#allow_headers NetworkServicesEdgeCacheService#allow_headers}
	AllowHeaders *[]*string `field:"optional" json:"allowHeaders" yaml:"allowHeaders"`
	// Specifies the content for the Access-Control-Allow-Methods response header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#allow_methods NetworkServicesEdgeCacheService#allow_methods}
	AllowMethods *[]*string `field:"optional" json:"allowMethods" yaml:"allowMethods"`
	// Specifies the list of origins that will be allowed to do CORS requests.
	//
	// This translates to the Access-Control-Allow-Origin response header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#allow_origins NetworkServicesEdgeCacheService#allow_origins}
	AllowOrigins *[]*string `field:"optional" json:"allowOrigins" yaml:"allowOrigins"`
	// If true, specifies the CORS policy is disabled.
	//
	// The default value is false, which indicates that the CORS policy is in effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#disabled NetworkServicesEdgeCacheService#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Specifies the content for the Access-Control-Allow-Headers response header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_services_edge_cache_service#expose_headers NetworkServicesEdgeCacheService#expose_headers}
	ExposeHeaders *[]*string `field:"optional" json:"exposeHeaders" yaml:"exposeHeaders"`
}

