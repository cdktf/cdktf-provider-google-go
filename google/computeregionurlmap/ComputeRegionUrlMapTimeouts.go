package computeregionurlmap


type ComputeRegionUrlMapTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#create ComputeRegionUrlMap#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#delete ComputeRegionUrlMap#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_url_map#update ComputeRegionUrlMap#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

