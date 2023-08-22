package computeregionurlmap


type ComputeRegionUrlMapTest struct {
	// Host portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#host ComputeRegionUrlMap#host}
	Host *string `field:"required" json:"host" yaml:"host"`
	// Path portion of the URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#path ComputeRegionUrlMap#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// A reference to expected RegionBackendService resource the given URL should be mapped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#service ComputeRegionUrlMap#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Description of this test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_url_map#description ComputeRegionUrlMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

