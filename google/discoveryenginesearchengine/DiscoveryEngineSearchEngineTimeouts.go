// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginesearchengine


type DiscoveryEngineSearchEngineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/discovery_engine_search_engine#create DiscoveryEngineSearchEngine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/discovery_engine_search_engine#delete DiscoveryEngineSearchEngine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/discovery_engine_search_engine#update DiscoveryEngineSearchEngine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

