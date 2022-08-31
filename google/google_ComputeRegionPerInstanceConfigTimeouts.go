// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeRegionPerInstanceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_per_instance_config#create ComputeRegionPerInstanceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_per_instance_config#delete ComputeRegionPerInstanceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_per_instance_config#update ComputeRegionPerInstanceConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

