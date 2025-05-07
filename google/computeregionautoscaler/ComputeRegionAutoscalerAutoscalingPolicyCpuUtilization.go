// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization struct {
	// The target CPU utilization that the autoscaler should maintain.
	//
	// Must be a float value in the range (0, 1]. If not specified, the
	// default is 0.6.
	//
	// If the CPU level is below the target utilization, the autoscaler
	// scales down the number of instances until it reaches the minimum
	// number of instances you specified or until the average CPU of
	// your instances reaches the target utilization.
	//
	// If the average CPU is above the target utilization, the autoscaler
	// scales up until it reaches the maximum number of instances you
	// specified or until the average utilization reaches the target
	// utilization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_region_autoscaler#target ComputeRegionAutoscaler#target}
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// Indicates whether predictive autoscaling based on CPU metric is enabled. Valid values are:.
	//
	// - NONE (default). No predictive method is used. The autoscaler scales the group to meet current demand based on real-time metrics.
	//
	// - OPTIMIZE_AVAILABILITY. Predictive autoscaling improves availability by monitoring daily and weekly load patterns and scaling out ahead of anticipated demand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/compute_region_autoscaler#predictive_method ComputeRegionAutoscaler#predictive_method}
	PredictiveMethod *string `field:"optional" json:"predictiveMethod" yaml:"predictiveMethod"`
}

