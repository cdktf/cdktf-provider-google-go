package computeregionhealthcheck


type ComputeRegionHealthCheckTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_health_check#create ComputeRegionHealthCheck#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_health_check#delete ComputeRegionHealthCheck#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_health_check#update ComputeRegionHealthCheck#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

