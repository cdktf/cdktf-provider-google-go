// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginerecommendationengine


type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfigRecommendedForYouConfig struct {
	// The type of event with which the engine is queried at prediction time.
	//
	// If set to 'generic', only 'view-item', 'media-play',and
	// 'media-complete' will be used as 'context-event' in engine training. If
	// set to 'view-home-page', 'view-home-page' will also be used as
	// 'context-events' in addition to 'view-item', 'media-play', and
	// 'media-complete'. Currently supported for the 'recommended-for-you'
	// engine. Currently supported values: 'view-home-page', 'generic'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#context_event_type DiscoveryEngineRecommendationEngine#context_event_type}
	ContextEventType *string `field:"optional" json:"contextEventType" yaml:"contextEventType"`
}

