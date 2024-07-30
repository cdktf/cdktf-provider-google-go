// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingmetric


type LoggingMetricBucketOptionsExponentialBuckets struct {
	// Must be greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/logging_metric#growth_factor LoggingMetric#growth_factor}
	GrowthFactor *float64 `field:"required" json:"growthFactor" yaml:"growthFactor"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/logging_metric#num_finite_buckets LoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"required" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/logging_metric#scale LoggingMetric#scale}
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
}

