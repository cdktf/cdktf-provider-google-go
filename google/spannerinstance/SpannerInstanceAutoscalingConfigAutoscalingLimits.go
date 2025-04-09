// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance


type SpannerInstanceAutoscalingConfigAutoscalingLimits struct {
	// Specifies maximum number of nodes allocated to the instance.
	//
	// If set, this number
	// should be greater than or equal to min_nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/spanner_instance#max_nodes SpannerInstance#max_nodes}
	MaxNodes *float64 `field:"optional" json:"maxNodes" yaml:"maxNodes"`
	// Specifies maximum number of processing units allocated to the instance.
	//
	// If set, this number should be multiples of 1000 and be greater than or equal to
	// min_processing_units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/spanner_instance#max_processing_units SpannerInstance#max_processing_units}
	MaxProcessingUnits *float64 `field:"optional" json:"maxProcessingUnits" yaml:"maxProcessingUnits"`
	// Specifies number of nodes allocated to the instance. If set, this number should be greater than or equal to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/spanner_instance#min_nodes SpannerInstance#min_nodes}
	MinNodes *float64 `field:"optional" json:"minNodes" yaml:"minNodes"`
	// Specifies minimum number of processing units allocated to the instance. If set, this number should be multiples of 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/spanner_instance#min_processing_units SpannerInstance#min_processing_units}
	MinProcessingUnits *float64 `field:"optional" json:"minProcessingUnits" yaml:"minProcessingUnits"`
}

