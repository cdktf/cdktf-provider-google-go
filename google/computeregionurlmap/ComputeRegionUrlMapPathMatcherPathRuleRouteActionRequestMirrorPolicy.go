package computeregionurlmap


type ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy struct {
	// The RegionBackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_url_map#backend_service ComputeRegionUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

