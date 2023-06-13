package computeregionbackendservice


type ComputeRegionBackendServiceConsistentHashHttpCookie struct {
	// Name of the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_backend_service#name ComputeRegionBackendService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Path to set for the cookie.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_backend_service#path ComputeRegionBackendService#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// ttl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_region_backend_service#ttl ComputeRegionBackendService#ttl}
	Ttl *ComputeRegionBackendServiceConsistentHashHttpCookieTtl `field:"optional" json:"ttl" yaml:"ttl"`
}

