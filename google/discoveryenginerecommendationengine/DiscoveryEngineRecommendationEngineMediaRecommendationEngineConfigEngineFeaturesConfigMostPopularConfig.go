// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginerecommendationengine


type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigMostPopularConfig struct {
	// The time window of which the engine is queried at training and prediction time.
	//
	// Positive integers only. The value translates to the
	// last X days of events. Currently required for the 'most-popular-items'
	// engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/discovery_engine_recommendation_engine#time_window_days DiscoveryEngineRecommendationEngine#time_window_days}
	TimeWindowDays *float64 `field:"optional" json:"timeWindowDays" yaml:"timeWindowDays"`
}

