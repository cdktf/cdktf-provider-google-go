package networkservicesedgecacheorigin


type NetworkServicesEdgeCacheOriginOriginRedirect struct {
	// The set of redirect response codes that the CDN follows. Values of [RedirectConditions](https://cloud.google.com/media-cdn/docs/reference/rest/v1/projects.locations.edgeCacheOrigins#redirectconditions) are accepted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/network_services_edge_cache_origin#redirect_conditions NetworkServicesEdgeCacheOrigin#redirect_conditions}
	RedirectConditions *[]*string `field:"optional" json:"redirectConditions" yaml:"redirectConditions"`
}

