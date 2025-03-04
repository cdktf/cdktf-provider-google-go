// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow struct {
	// Auth URL for Authorization Code Flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/integration_connectors_connection#auth_uri IntegrationConnectorsConnection#auth_uri}
	AuthUri *string `field:"optional" json:"authUri" yaml:"authUri"`
	// Client ID for user-provided OAuth app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/integration_connectors_connection#client_id IntegrationConnectorsConnection#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/integration_connectors_connection#client_secret IntegrationConnectorsConnection#client_secret}
	ClientSecret *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Whether to enable PKCE when the user performs the auth code flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/integration_connectors_connection#enable_pkce IntegrationConnectorsConnection#enable_pkce}
	EnablePkce interface{} `field:"optional" json:"enablePkce" yaml:"enablePkce"`
	// Scopes the connection will request when the user performs the auth code flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/integration_connectors_connection#scopes IntegrationConnectorsConnection#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}

