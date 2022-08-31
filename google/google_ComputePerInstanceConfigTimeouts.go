// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputePerInstanceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_per_instance_config#create ComputePerInstanceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_per_instance_config#delete ComputePerInstanceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_per_instance_config#update ComputePerInstanceConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

