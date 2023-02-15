package computeaddress


type ComputeAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_address#create ComputeAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_address#delete ComputeAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

