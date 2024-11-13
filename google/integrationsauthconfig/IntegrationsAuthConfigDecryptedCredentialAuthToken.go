// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialAuthToken struct {
	// The token for the auth type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integrations_auth_config#token IntegrationsAuthConfig#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Authentication type, e.g. "Basic", "Bearer", etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/integrations_auth_config#type IntegrationsAuthConfig#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

