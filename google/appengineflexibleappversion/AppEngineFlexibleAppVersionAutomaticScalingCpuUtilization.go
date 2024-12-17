// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion


type AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization struct {
	// Target CPU utilization ratio to maintain when scaling. Must be between 0 and 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/app_engine_flexible_app_version#target_utilization AppEngineFlexibleAppVersion#target_utilization}
	TargetUtilization *float64 `field:"required" json:"targetUtilization" yaml:"targetUtilization"`
	// Period of time over which CPU utilization is calculated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/app_engine_flexible_app_version#aggregation_window_length AppEngineFlexibleAppVersion#aggregation_window_length}
	AggregationWindowLength *string `field:"optional" json:"aggregationWindowLength" yaml:"aggregationWindowLength"`
}

