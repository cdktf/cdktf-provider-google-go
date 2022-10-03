package computeinstancetemplate


type ComputeInstanceTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#create ComputeInstanceTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#delete ComputeInstanceTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

