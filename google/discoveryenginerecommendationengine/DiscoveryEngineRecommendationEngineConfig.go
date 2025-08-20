// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginerecommendationengine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineRecommendationEngineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The data stores associated with this engine.
	//
	// For SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#data_store_ids DiscoveryEngineRecommendationEngine#data_store_ids}
	DataStoreIds *[]*string `field:"required" json:"dataStoreIds" yaml:"dataStoreIds"`
	// Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#display_name DiscoveryEngineRecommendationEngine#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Unique ID to use for Recommendation Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#engine_id DiscoveryEngineRecommendationEngine#engine_id}
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// The geographic location where the data store should reside. The value can only be one of "global", "us" and "eu".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#location DiscoveryEngineRecommendationEngine#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// common_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#common_config DiscoveryEngineRecommendationEngine#common_config}
	CommonConfig *DiscoveryEngineRecommendationEngineCommonConfig `field:"optional" json:"commonConfig" yaml:"commonConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#id DiscoveryEngineRecommendationEngine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The industry vertical that the engine registers.
	//
	// The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine. Default value: "GENERIC" Possible values: ["GENERIC", "MEDIA"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#industry_vertical DiscoveryEngineRecommendationEngine#industry_vertical}
	IndustryVertical *string `field:"optional" json:"industryVertical" yaml:"industryVertical"`
	// media_recommendation_engine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#media_recommendation_engine_config DiscoveryEngineRecommendationEngine#media_recommendation_engine_config}
	MediaRecommendationEngineConfig *DiscoveryEngineRecommendationEngineMediaRecommendationEngineConfig `field:"optional" json:"mediaRecommendationEngineConfig" yaml:"mediaRecommendationEngineConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#project DiscoveryEngineRecommendationEngine#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/discovery_engine_recommendation_engine#timeouts DiscoveryEngineRecommendationEngine#timeouts}
	Timeouts *DiscoveryEngineRecommendationEngineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

