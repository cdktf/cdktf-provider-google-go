package computeinstancetemplate


type ComputeInstanceTemplateSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#key ComputeInstanceTemplate#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#operator ComputeInstanceTemplate#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#values ComputeInstanceTemplate#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

