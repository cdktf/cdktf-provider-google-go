package computetargetinstance


type ComputeTargetInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_target_instance#create ComputeTargetInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_target_instance#delete ComputeTargetInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

