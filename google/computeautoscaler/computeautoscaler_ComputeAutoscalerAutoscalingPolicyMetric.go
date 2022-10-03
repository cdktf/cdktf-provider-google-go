package computeautoscaler


type ComputeAutoscalerAutoscalingPolicyMetric struct {
	// The identifier (type) of the Stackdriver Monitoring metric. The metric cannot have negative values.
	//
	// The metric must have a value type of INT64 or DOUBLE.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#name ComputeAutoscaler#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target value of the metric that autoscaler should maintain.
	//
	// This must be a positive value. A utilization
	// metric scales number of virtual machines handling requests
	// to increase or decrease proportionally to the metric.
	//
	// For example, a good metric to use as a utilizationTarget is
	// www.googleapis.com/compute/instance/network/received_bytes_count.
	// The autoscaler will work to keep this value constant for each
	// of the instances.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#target ComputeAutoscaler#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
	// Defines how target utilization value is expressed for a Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_autoscaler#type ComputeAutoscaler#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

