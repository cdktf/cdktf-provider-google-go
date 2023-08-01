package computeautoscaler


type ComputeAutoscalerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_autoscaler#create ComputeAutoscaler#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_autoscaler#delete ComputeAutoscaler#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_autoscaler#update ComputeAutoscaler#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

