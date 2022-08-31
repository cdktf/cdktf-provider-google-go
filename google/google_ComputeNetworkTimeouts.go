// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeNetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network#create ComputeNetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network#delete ComputeNetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network#update ComputeNetwork#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

