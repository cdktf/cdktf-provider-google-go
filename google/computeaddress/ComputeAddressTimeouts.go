package computeaddress


type ComputeAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_address#create ComputeAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_address#delete ComputeAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

