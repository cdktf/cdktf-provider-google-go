package computeregionurlmap


type ComputeRegionUrlMapPathMatcherRouteRulesMatchRulesMetadataFiltersFilterLabels struct {
	// Name of metadata label.
	//
	// The name can have a maximum length of 1024 characters
	// and must be at least 1 character long.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_url_map#name ComputeRegionUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the label must match the specified value. value can have a maximum length of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_url_map#value ComputeRegionUrlMap#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

