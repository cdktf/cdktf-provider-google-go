// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig


type IntegrationsAuthConfigDecryptedCredentialJwt struct {
	// Identifies which algorithm is used to generate the signature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/integrations_auth_config#jwt_header IntegrationsAuthConfig#jwt_header}
	JwtHeader *string `field:"optional" json:"jwtHeader" yaml:"jwtHeader"`
	// Contains a set of claims.
	//
	// The JWT specification defines seven Registered Claim Names which are the standard fields commonly included in tokens. Custom claims are usually also included, depending on the purpose of the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/integrations_auth_config#jwt_payload IntegrationsAuthConfig#jwt_payload}
	JwtPayload *string `field:"optional" json:"jwtPayload" yaml:"jwtPayload"`
	// User's pre-shared secret to sign the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/integrations_auth_config#secret IntegrationsAuthConfig#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

