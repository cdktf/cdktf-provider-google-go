// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginesearchengine


type DiscoveryEngineSearchEngineCommonConfig struct {
	// The name of the company, business or entity that is associated with the engine.
	//
	// Setting this may help improve LLM related features.cd
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/discovery_engine_search_engine#company_name DiscoveryEngineSearchEngine#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
}

