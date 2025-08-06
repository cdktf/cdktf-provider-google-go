// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection


type IntegrationConnectorsConnectionAuthConfig struct {
	// authType of the Connection Possible values: ["AUTH_TYPE_UNSPECIFIED", "USER_PASSWORD", "OAUTH2_JWT_BEARER", "OAUTH2_CLIENT_CREDENTIALS", "SSH_PUBLIC_KEY", "OAUTH2_AUTH_CODE_FLOW"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#auth_type IntegrationConnectorsConnection#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// additional_variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#additional_variable IntegrationConnectorsConnection#additional_variable}
	AdditionalVariable interface{} `field:"optional" json:"additionalVariable" yaml:"additionalVariable"`
	// The type of authentication configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#auth_key IntegrationConnectorsConnection#auth_key}
	AuthKey *string `field:"optional" json:"authKey" yaml:"authKey"`
	// oauth2_auth_code_flow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#oauth2_auth_code_flow IntegrationConnectorsConnection#oauth2_auth_code_flow}
	Oauth2AuthCodeFlow *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow `field:"optional" json:"oauth2AuthCodeFlow" yaml:"oauth2AuthCodeFlow"`
	// oauth2_client_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#oauth2_client_credentials IntegrationConnectorsConnection#oauth2_client_credentials}
	Oauth2ClientCredentials *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials `field:"optional" json:"oauth2ClientCredentials" yaml:"oauth2ClientCredentials"`
	// oauth2_jwt_bearer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#oauth2_jwt_bearer IntegrationConnectorsConnection#oauth2_jwt_bearer}
	Oauth2JwtBearer *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer `field:"optional" json:"oauth2JwtBearer" yaml:"oauth2JwtBearer"`
	// ssh_public_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#ssh_public_key IntegrationConnectorsConnection#ssh_public_key}
	SshPublicKey *IntegrationConnectorsConnectionAuthConfigSshPublicKey `field:"optional" json:"sshPublicKey" yaml:"sshPublicKey"`
	// user_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/integration_connectors_connection#user_password IntegrationConnectorsConnection#user_password}
	UserPassword *IntegrationConnectorsConnectionAuthConfigUserPassword `field:"optional" json:"userPassword" yaml:"userPassword"`
}

