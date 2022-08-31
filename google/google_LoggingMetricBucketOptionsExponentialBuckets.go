// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type LoggingMetricBucketOptionsExponentialBuckets struct {
	// Must be greater than 1.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#growth_factor LoggingMetric#growth_factor}
	GrowthFactor *float64 `field:"optional" json:"growthFactor" yaml:"growthFactor"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#num_finite_buckets LoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"optional" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/logging_metric#scale LoggingMetric#scale}
	Scale *float64 `field:"optional" json:"scale" yaml:"scale"`
}

