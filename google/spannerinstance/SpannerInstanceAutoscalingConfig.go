// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance


type SpannerInstanceAutoscalingConfig struct {
	// asymmetric_autoscaling_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/spanner_instance#asymmetric_autoscaling_options SpannerInstance#asymmetric_autoscaling_options}
	AsymmetricAutoscalingOptions interface{} `field:"optional" json:"asymmetricAutoscalingOptions" yaml:"asymmetricAutoscalingOptions"`
	// autoscaling_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/spanner_instance#autoscaling_limits SpannerInstance#autoscaling_limits}
	AutoscalingLimits *SpannerInstanceAutoscalingConfigAutoscalingLimits `field:"optional" json:"autoscalingLimits" yaml:"autoscalingLimits"`
	// autoscaling_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/spanner_instance#autoscaling_targets SpannerInstance#autoscaling_targets}
	AutoscalingTargets *SpannerInstanceAutoscalingConfigAutoscalingTargets `field:"optional" json:"autoscalingTargets" yaml:"autoscalingTargets"`
}

