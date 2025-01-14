// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance


type SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits struct {
	// The maximum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/spanner_instance#max_nodes SpannerInstance#max_nodes}
	MaxNodes *float64 `field:"required" json:"maxNodes" yaml:"maxNodes"`
	// The minimum number of nodes for this specific replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/spanner_instance#min_nodes SpannerInstance#min_nodes}
	MinNodes *float64 `field:"required" json:"minNodes" yaml:"minNodes"`
}

