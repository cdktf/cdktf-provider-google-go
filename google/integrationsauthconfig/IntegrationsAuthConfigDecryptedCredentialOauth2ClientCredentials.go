// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials struct {
	// The client's ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#client_id IntegrationsAuthConfig#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client's secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#client_secret IntegrationsAuthConfig#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Represent how to pass parameters to fetch access token Possible values: ["REQUEST_TYPE_UNSPECIFIED", "REQUEST_BODY", "QUERY_PARAMETERS", "ENCODED_HEADER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#request_type IntegrationsAuthConfig#request_type}
	RequestType *string `field:"optional" json:"requestType" yaml:"requestType"`
	// A space-delimited list of requested scope permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#scope IntegrationsAuthConfig#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The token endpoint is used by the client to obtain an access token by presenting its authorization grant or refresh token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#token_endpoint IntegrationsAuthConfig#token_endpoint}
	TokenEndpoint *string `field:"optional" json:"tokenEndpoint" yaml:"tokenEndpoint"`
	// token_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/integrations_auth_config#token_params IntegrationsAuthConfig#token_params}
	TokenParams *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsTokenParams `field:"optional" json:"tokenParams" yaml:"tokenParams"`
}

