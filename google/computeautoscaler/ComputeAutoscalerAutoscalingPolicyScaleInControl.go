package computeautoscaler


type ComputeAutoscalerAutoscalingPolicyScaleInControl struct {
	// max_scaled_in_replicas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#max_scaled_in_replicas ComputeAutoscaler#max_scaled_in_replicas}
	MaxScaledInReplicas *ComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas `field:"optional" json:"maxScaledInReplicas" yaml:"maxScaledInReplicas"`
	// How long back autoscaling should look when computing recommendations to include directives regarding slower scale down, as described above.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#time_window_sec ComputeAutoscaler#time_window_sec}
	TimeWindowSec *float64 `field:"optional" json:"timeWindowSec" yaml:"timeWindowSec"`
}

