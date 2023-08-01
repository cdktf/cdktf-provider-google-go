package loggingmetric


type LoggingMetricBucketOptionsLinearBuckets struct {
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/logging_metric#num_finite_buckets LoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"optional" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Lower bound of the first bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/logging_metric#offset LoggingMetric#offset}
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/logging_metric#width LoggingMetric#width}
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

