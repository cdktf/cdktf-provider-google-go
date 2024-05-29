// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformoauthidpconfig


type IdentityPlatformOauthIdpConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/identity_platform_oauth_idp_config#create IdentityPlatformOauthIdpConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/identity_platform_oauth_idp_config#delete IdentityPlatformOauthIdpConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/identity_platform_oauth_idp_config#update IdentityPlatformOauthIdpConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

