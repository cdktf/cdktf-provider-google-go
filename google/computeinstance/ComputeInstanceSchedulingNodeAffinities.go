package computeinstance


type ComputeInstanceSchedulingNodeAffinities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#key ComputeInstance#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#operator ComputeInstance#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#values ComputeInstance#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

