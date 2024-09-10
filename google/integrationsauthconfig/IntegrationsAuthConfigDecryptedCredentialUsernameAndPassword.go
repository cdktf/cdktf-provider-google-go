// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword struct {
	// Password to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/integrations_auth_config#password IntegrationsAuthConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Username to be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/integrations_auth_config#username IntegrationsAuthConfig#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

