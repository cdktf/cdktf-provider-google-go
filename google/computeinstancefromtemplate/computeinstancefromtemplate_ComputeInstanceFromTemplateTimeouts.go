package computeinstancefromtemplate


type ComputeInstanceFromTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#create ComputeInstanceFromTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#delete ComputeInstanceFromTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#update ComputeInstanceFromTemplate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

