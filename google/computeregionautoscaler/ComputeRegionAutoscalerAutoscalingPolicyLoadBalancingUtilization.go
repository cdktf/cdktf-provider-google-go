package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization struct {
	// Fraction of backend capacity utilization (set in HTTP(s) load balancing configuration) that autoscaler should maintain.
	//
	// Must
	// be a positive float value. If not defined, the default is 0.8.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_autoscaler#target ComputeRegionAutoscaler#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
}

