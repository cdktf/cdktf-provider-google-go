package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionUrlRewrite struct {
	// Before forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite.
	//
	// The value must be from 1 to 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_url_map#host_rewrite ComputeRegionUrlMap#host_rewrite}
	HostRewrite *string `field:"optional" json:"hostRewrite" yaml:"hostRewrite"`
	// Before forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite.
	//
	// The value must be from 1 to 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_url_map#path_prefix_rewrite ComputeRegionUrlMap#path_prefix_rewrite}
	PathPrefixRewrite *string `field:"optional" json:"pathPrefixRewrite" yaml:"pathPrefixRewrite"`
}

