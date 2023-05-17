package computetargetinstance


type ComputeTargetInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_target_instance#create ComputeTargetInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_target_instance#delete ComputeTargetInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

