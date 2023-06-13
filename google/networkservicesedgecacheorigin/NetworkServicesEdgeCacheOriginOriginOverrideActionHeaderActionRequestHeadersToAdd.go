package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginOriginOverrideActionHeaderActionRequestHeadersToAdd struct {
	// The name of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/network_services_edge_cache_origin#header_name NetworkServicesEdgeCacheOrigin#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// The value of the header to add.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/network_services_edge_cache_origin#header_value NetworkServicesEdgeCacheOrigin#header_value}
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
	// Whether to replace all existing headers with the same name.
	//
	// By default, added header values are appended
	// to the response or request headers with the
	// same field names. The added values are
	// separated by commas.
	//
	// To overwrite existing values, set 'replace' to 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/network_services_edge_cache_origin#replace NetworkServicesEdgeCacheOrigin#replace}
	Replace interface{} `field:"optional" json:"replace" yaml:"replace"`
}

