package computeglobaladdress


type ComputeGlobalAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_global_address#create ComputeGlobalAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_global_address#delete ComputeGlobalAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

