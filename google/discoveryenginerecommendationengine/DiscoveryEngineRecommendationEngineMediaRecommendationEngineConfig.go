// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginerecommendationengine


type DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig struct {
	// engine_features_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_recommendation_engine#engine_features_config DiscoveryEngineRecommendationEngine#engine_features_config}
	EngineFeaturesConfig *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigEngineFeaturesConfig `field:"optional" json:"engineFeaturesConfig" yaml:"engineFeaturesConfig"`
	// The optimization objective.
	//
	// e.g., 'cvr'.
	// This field together with MediaRecommendationEngineConfig.type describes
	// engine metadata to use to control engine training and serving.
	// Currently supported values: 'ctr', 'cvr'.
	// If not specified, we choose default based on engine type. Default depends on type of recommendation:
	// 'recommended-for-you' => 'ctr'
	// 'others-you-may-like' => 'ctr'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_recommendation_engine#optimization_objective DiscoveryEngineRecommendationEngine#optimization_objective}
	OptimizationObjective *string `field:"optional" json:"optimizationObjective" yaml:"optimizationObjective"`
	// optimization_objective_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_recommendation_engine#optimization_objective_config DiscoveryEngineRecommendationEngine#optimization_objective_config}
	OptimizationObjectiveConfig *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfigOptimizationObjectiveConfig `field:"optional" json:"optimizationObjectiveConfig" yaml:"optimizationObjectiveConfig"`
	// The training state that the engine is in (e.g. 'TRAINING' or 'PAUSED'). Since part of the cost of running the service is frequency of training - this can be used to determine when to train engine in order to control cost. If not specified: the default value for 'CreateEngine' method is 'TRAINING'. The default value for 'UpdateEngine' method is to keep the state the same as before. Possible values: ["PAUSED", "TRAINING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_recommendation_engine#training_state DiscoveryEngineRecommendationEngine#training_state}
	TrainingState *string `field:"optional" json:"trainingState" yaml:"trainingState"`
	// The type of engine.
	//
	// e.g., 'recommended-for-you'.
	// This field together with MediaRecommendationEngineConfig.optimizationObjective describes
	// engine metadata to use to control engine training and serving.
	// Currently supported values: 'recommended-for-you', 'others-you-may-like',
	// 'more-like-this', 'most-popular-items'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/discovery_engine_recommendation_engine#type DiscoveryEngineRecommendationEngine#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

