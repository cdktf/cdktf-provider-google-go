package computeregionurlmap


type ComputeRegionUrlMapPathMatcherPathRuleRouteActionRequestMirrorPolicy struct {
	// The RegionBackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#backend_service ComputeRegionUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

