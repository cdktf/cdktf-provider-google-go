package computeinstancefromtemplate


type ComputeInstanceFromTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_from_template#key ComputeInstanceFromTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_from_template#operator ComputeInstanceFromTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_instance_from_template#values ComputeInstanceFromTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

