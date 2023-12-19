// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance


type SpannerInstanceAutoscalingConfigAutoscalingTargets struct {
	// Specifies the target high priority cpu utilization percentage that the autoscaler should be trying to achieve for the instance.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization)..
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/spanner_instance#high_priority_cpu_utilization_percent SpannerInstance#high_priority_cpu_utilization_percent}
	HighPriorityCpuUtilizationPercent *float64 `field:"optional" json:"highPriorityCpuUtilizationPercent" yaml:"highPriorityCpuUtilizationPercent"`
	// Specifies the target storage utilization percentage that the autoscaler should be trying to achieve for the instance.
	//
	// This number is on a scale from 0 (no utilization) to 100 (full utilization).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/spanner_instance#storage_utilization_percent SpannerInstance#storage_utilization_percent}
	StorageUtilizationPercent *float64 `field:"optional" json:"storageUtilizationPercent" yaml:"storageUtilizationPercent"`
}

