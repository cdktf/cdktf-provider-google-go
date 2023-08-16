package computeregionsslpolicy


type ComputeRegionSslPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_ssl_policy#create ComputeRegionSslPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_ssl_policy#delete ComputeRegionSslPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_ssl_policy#update ComputeRegionSslPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

