package computeinstancetemplate


type ComputeInstanceTemplateGuestAccelerator struct {
	// The number of the guest accelerator cards exposed to this instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#count ComputeInstanceTemplate#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource to expose to this instance. E.g. nvidia-tesla-k80.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#type ComputeInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}
