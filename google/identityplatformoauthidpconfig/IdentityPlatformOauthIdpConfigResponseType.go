// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformoauthidpconfig


type IdentityPlatformOauthIdpConfigResponseType struct {
	// If true, authorization code is returned from IdP's authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/identity_platform_oauth_idp_config#code IdentityPlatformOauthIdpConfig#code}
	Code interface{} `field:"optional" json:"code" yaml:"code"`
	// If true, ID token is returned from IdP's authorization endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/identity_platform_oauth_idp_config#id_token IdentityPlatformOauthIdpConfig#id_token}
	IdToken interface{} `field:"optional" json:"idToken" yaml:"idToken"`
}

