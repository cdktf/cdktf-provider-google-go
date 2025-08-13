// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials struct {
	// Secret version of Password for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/integration_connectors_connection#client_id IntegrationConnectorsConnection#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// client_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/integration_connectors_connection#client_secret IntegrationConnectorsConnection#client_secret}
	ClientSecret *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsClientSecret `field:"optional" json:"clientSecret" yaml:"clientSecret"`
}

