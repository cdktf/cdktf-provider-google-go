package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicyScaleInControl struct {
	// max_scaled_in_replicas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#max_scaled_in_replicas ComputeRegionAutoscaler#max_scaled_in_replicas}
	MaxScaledInReplicas *ComputeRegionAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `field:"optional" json:"maxScaledInReplicas" yaml:"maxScaledInReplicas"`
	// How long back autoscaling should look when computing recommendations to include directives regarding slower scale down, as described above.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#time_window_sec ComputeRegionAutoscaler#time_window_sec}
	TimeWindowSec *float64 `field:"optional" json:"timeWindowSec" yaml:"timeWindowSec"`
}

