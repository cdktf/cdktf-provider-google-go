// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtool


type DialogflowCxToolOpenApiSpecAuthentication struct {
	// api_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_tool#api_key_config DialogflowCxTool#api_key_config}
	ApiKeyConfig *DialogflowCxToolOpenApiSpecAuthenticationApiKeyConfig `field:"optional" json:"apiKeyConfig" yaml:"apiKeyConfig"`
	// bearer_token_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_tool#bearer_token_config DialogflowCxTool#bearer_token_config}
	BearerTokenConfig *DialogflowCxToolOpenApiSpecAuthenticationBearerTokenConfig `field:"optional" json:"bearerTokenConfig" yaml:"bearerTokenConfig"`
	// oauth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_tool#oauth_config DialogflowCxTool#oauth_config}
	OauthConfig *DialogflowCxToolOpenApiSpecAuthenticationOauthConfig `field:"optional" json:"oauthConfig" yaml:"oauthConfig"`
	// service_agent_auth_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/dialogflow_cx_tool#service_agent_auth_config DialogflowCxTool#service_agent_auth_config}
	ServiceAgentAuthConfig *DialogflowCxToolOpenApiSpecAuthenticationServiceAgentAuthConfig `field:"optional" json:"serviceAgentAuthConfig" yaml:"serviceAgentAuthConfig"`
}

