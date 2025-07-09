// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings


type DialogflowCxGenerativeSettingsKnowledgeConnectorSettings struct {
	// Name of the virtual agent. Used for LLM prompt. Can be left empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#agent DialogflowCxGenerativeSettings#agent}
	Agent *string `field:"optional" json:"agent" yaml:"agent"`
	// Identity of the agent, e.g. "virtual agent", "AI assistant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#agent_identity DialogflowCxGenerativeSettings#agent_identity}
	AgentIdentity *string `field:"optional" json:"agentIdentity" yaml:"agentIdentity"`
	// Agent scope, e.g. "Example company website", "internal Example company website for employees", "manual of car owner".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#agent_scope DialogflowCxGenerativeSettings#agent_scope}
	AgentScope *string `field:"optional" json:"agentScope" yaml:"agentScope"`
	// Name of the company, organization or other entity that the agent represents.
	//
	// Used for knowledge connector LLM prompt and for knowledge search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#business DialogflowCxGenerativeSettings#business}
	Business *string `field:"optional" json:"business" yaml:"business"`
	// Company description, used for LLM prompt, e.g. "a family company selling freshly roasted coffee beans".''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#business_description DialogflowCxGenerativeSettings#business_description}
	BusinessDescription *string `field:"optional" json:"businessDescription" yaml:"businessDescription"`
	// Whether to disable fallback to Data Store search results (in case the LLM couldn't pick a proper answer).
	//
	// Per default the feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/dialogflow_cx_generative_settings#disable_data_store_fallback DialogflowCxGenerativeSettings#disable_data_store_fallback}
	DisableDataStoreFallback interface{} `field:"optional" json:"disableDataStoreFallback" yaml:"disableDataStoreFallback"`
}

