package computeregionurlmap


type ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch struct {
	// The end of the range (exclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#range_end ComputeRegionUrlMap#range_end}
	RangeEnd *float64 `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// The start of the range (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_region_url_map#range_start ComputeRegionUrlMap#range_start}
	RangeStart *float64 `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

