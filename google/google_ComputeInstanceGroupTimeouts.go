// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeInstanceGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_group#create ComputeInstanceGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_group#delete ComputeInstanceGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_group#update ComputeInstanceGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

