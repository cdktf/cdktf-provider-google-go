// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginechatengine


type DiscoveryEngineChatEngineCommonConfig struct {
	// The name of the company, business or entity that is associated with the engine.
	//
	// Setting this may help improve LLM related features.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/discovery_engine_chat_engine#company_name DiscoveryEngineChatEngine#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
}

