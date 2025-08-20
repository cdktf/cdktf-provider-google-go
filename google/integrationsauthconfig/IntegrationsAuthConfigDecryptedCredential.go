// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredential struct {
	// Credential type associated with auth configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#credential_type IntegrationsAuthConfig#credential_type}
	CredentialType *string `field:"required" json:"credentialType" yaml:"credentialType"`
	// auth_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#auth_token IntegrationsAuthConfig#auth_token}
	AuthToken *IntegrationsAuthConfigDecryptedCredentialAuthToken `field:"optional" json:"authToken" yaml:"authToken"`
	// jwt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#jwt IntegrationsAuthConfig#jwt}
	Jwt *IntegrationsAuthConfigDecryptedCredentialJwt `field:"optional" json:"jwt" yaml:"jwt"`
	// oauth2_authorization_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#oauth2_authorization_code IntegrationsAuthConfig#oauth2_authorization_code}
	Oauth2AuthorizationCode *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode `field:"optional" json:"oauth2AuthorizationCode" yaml:"oauth2AuthorizationCode"`
	// oauth2_client_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#oauth2_client_credentials IntegrationsAuthConfig#oauth2_client_credentials}
	Oauth2ClientCredentials *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials `field:"optional" json:"oauth2ClientCredentials" yaml:"oauth2ClientCredentials"`
	// oidc_token block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#oidc_token IntegrationsAuthConfig#oidc_token}
	OidcToken *IntegrationsAuthConfigDecryptedCredentialOidcToken `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// service_account_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#service_account_credentials IntegrationsAuthConfig#service_account_credentials}
	ServiceAccountCredentials *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials `field:"optional" json:"serviceAccountCredentials" yaml:"serviceAccountCredentials"`
	// username_and_password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/integrations_auth_config#username_and_password IntegrationsAuthConfig#username_and_password}
	UsernameAndPassword *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword `field:"optional" json:"usernameAndPassword" yaml:"usernameAndPassword"`
}

