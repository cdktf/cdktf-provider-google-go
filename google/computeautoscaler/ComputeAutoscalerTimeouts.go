package computeautoscaler


type ComputeAutoscalerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#create ComputeAutoscaler#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#delete ComputeAutoscaler#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#update ComputeAutoscaler#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

