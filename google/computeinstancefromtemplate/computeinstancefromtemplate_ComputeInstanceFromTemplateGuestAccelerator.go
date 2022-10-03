package computeinstancefromtemplate


type ComputeInstanceFromTemplateGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#count ComputeInstanceFromTemplate#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#type ComputeInstanceFromTemplate#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

