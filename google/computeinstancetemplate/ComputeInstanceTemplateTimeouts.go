package computeinstancetemplate


type ComputeInstanceTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_template#create ComputeInstanceTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_template#delete ComputeInstanceTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

