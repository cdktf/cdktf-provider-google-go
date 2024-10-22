// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialOidcToken struct {
	// Audience to be used when generating OIDC token.
	//
	// The audience claim identifies the recipients that the JWT is intended for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/integrations_auth_config#audience IntegrationsAuthConfig#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
	// The service account email to be used as the identity for the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/integrations_auth_config#service_account_email IntegrationsAuthConfig#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
}

