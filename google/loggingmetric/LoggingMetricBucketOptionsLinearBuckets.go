// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingmetric


type LoggingMetricBucketOptionsLinearBuckets struct {
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/logging_metric#num_finite_buckets LoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"required" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Lower bound of the first bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/logging_metric#offset LoggingMetric#offset}
	Offset *float64 `field:"required" json:"offset" yaml:"offset"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/logging_metric#width LoggingMetric#width}
	Width *float64 `field:"required" json:"width" yaml:"width"`
}

