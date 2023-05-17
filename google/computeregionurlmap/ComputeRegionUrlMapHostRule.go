package computeregionurlmap


type ComputeRegionUrlMapHostRule struct {
	// The list of host patterns to match.
	//
	// They must be valid
	// hostnames, except * will match any string of ([a-z0-9-.]*). In
	// that case, * must be the first character and must be followed in
	// the pattern by either - or ..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_url_map#hosts ComputeRegionUrlMap#hosts}
	Hosts *[]*string `field:"required" json:"hosts" yaml:"hosts"`
	// The name of the PathMatcher to use to match the path portion of the URL if the hostRule matches the URL's host portion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_url_map#path_matcher ComputeRegionUrlMap#path_matcher}
	PathMatcher *string `field:"required" json:"pathMatcher" yaml:"pathMatcher"`
	// An optional description of this HostRule. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_url_map#description ComputeRegionUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

