package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicyMetric struct {
	// The identifier (type) of the Stackdriver Monitoring metric. The metric cannot have negative values.
	//
	// The metric must have a value type of INT64 or DOUBLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.0/docs/resources/compute_region_autoscaler#name ComputeRegionAutoscaler#name}
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.0/docs/resources/compute_region_autoscaler#target ComputeRegionAutoscaler#target}
	Target *float64 `field:"optional" json:"target" yaml:"target"`
	// Defines how target utilization value is expressed for a Stackdriver Monitoring metric. Possible values: ["GAUGE", "DELTA_PER_SECOND", "DELTA_PER_MINUTE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.72.0/docs/resources/compute_region_autoscaler#type ComputeRegionAutoscaler#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

