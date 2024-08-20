// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginesearchengine


type DiscoveryEngineSearchEngineSearchEngineConfig struct {
	// The add-on that this search engine enables. Possible values: ["SEARCH_ADD_ON_LLM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/discovery_engine_search_engine#search_add_ons DiscoveryEngineSearchEngine#search_add_ons}
	SearchAddOns *[]*string `field:"optional" json:"searchAddOns" yaml:"searchAddOns"`
	// The search feature tier of this engine.
	//
	// Defaults to SearchTier.SEARCH_TIER_STANDARD if not specified. Default value: "SEARCH_TIER_STANDARD" Possible values: ["SEARCH_TIER_STANDARD", "SEARCH_TIER_ENTERPRISE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.42.0/docs/resources/discovery_engine_search_engine#search_tier DiscoveryEngineSearchEngine#search_tier}
	SearchTier *string `field:"optional" json:"searchTier" yaml:"searchTier"`
}

