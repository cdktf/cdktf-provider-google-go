// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginerecommendationengine


type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig struct {
	// The name of the field to target. Currently supported values: 'watch-percentage', 'watch-time'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_recommendation_engine#target_field DiscoveryEngineRecommendationEngine#target_field}
	TargetField *string `field:"optional" json:"targetField" yaml:"targetField"`
	// The threshold to be applied to the target (e.g., 0.5).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/discovery_engine_recommendation_engine#target_field_value_float DiscoveryEngineRecommendationEngine#target_field_value_float}
	TargetFieldValueFloat *float64 `field:"optional" json:"targetFieldValueFloat" yaml:"targetFieldValueFloat"`
}

