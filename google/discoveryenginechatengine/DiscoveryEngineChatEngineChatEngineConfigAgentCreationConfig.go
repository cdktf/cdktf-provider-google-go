// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginechatengine


type DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig struct {
	// The default language of the agent as a language tag.
	//
	// See [Language Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/discovery_engine_chat_engine#default_language_code DiscoveryEngineChatEngine#default_language_code}
	DefaultLanguageCode *string `field:"required" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/discovery_engine_chat_engine#time_zone DiscoveryEngineChatEngine#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Name of the company, organization or other entity that the agent represents.
	//
	// Used for knowledge connector LLM prompt and for knowledge search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/discovery_engine_chat_engine#business DiscoveryEngineChatEngine#business}
	Business *string `field:"optional" json:"business" yaml:"business"`
	// Agent location for Agent creation, currently supported values: global/us/eu, it needs to be the same region as the Chat Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/discovery_engine_chat_engine#location DiscoveryEngineChatEngine#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
}

