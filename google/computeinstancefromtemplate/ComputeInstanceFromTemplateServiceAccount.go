package computeinstancefromtemplate


type ComputeInstanceFromTemplateServiceAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#email ComputeInstanceFromTemplate#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_from_template#scopes ComputeInstanceFromTemplate#scopes}.
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

