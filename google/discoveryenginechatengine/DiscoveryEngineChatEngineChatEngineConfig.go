// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginechatengine


type DiscoveryEngineChatEngineChatEngineConfig struct {
	// agent_creation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_chat_engine#agent_creation_config DiscoveryEngineChatEngine#agent_creation_config}
	AgentCreationConfig *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig `field:"optional" json:"agentCreationConfig" yaml:"agentCreationConfig"`
	// If the flag set to true, we allow the agent and engine are in different locations, otherwise the agent and engine are required to be in the same location.
	//
	// The flag is set to false by default.
	// Note that the 'allow_cross_region' are one-time consumed by and passed
	// to EngineService.CreateEngine. It means they cannot be retrieved using
	// EngineService.GetEngine or EngineService.ListEngines API after engine
	// creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_chat_engine#allow_cross_region DiscoveryEngineChatEngine#allow_cross_region}
	AllowCrossRegion interface{} `field:"optional" json:"allowCrossRegion" yaml:"allowCrossRegion"`
	// The resource name of an existing Dialogflow agent to link to this Chat Engine.
	//
	// Format: 'projects/<Project_ID>/locations/<Location_ID>/agents/<Agent_ID>'.
	// Exactly one of 'agent_creation_config' or 'dialogflow_agent_to_link' must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/discovery_engine_chat_engine#dialogflow_agent_to_link DiscoveryEngineChatEngine#dialogflow_agent_to_link}
	DialogflowAgentToLink *string `field:"optional" json:"dialogflowAgentToLink" yaml:"dialogflowAgentToLink"`
}

