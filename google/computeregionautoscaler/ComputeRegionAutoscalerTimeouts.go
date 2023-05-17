package computeregionautoscaler


type ComputeRegionAutoscalerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_autoscaler#create ComputeRegionAutoscaler#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_autoscaler#delete ComputeRegionAutoscaler#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_region_autoscaler#update ComputeRegionAutoscaler#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

