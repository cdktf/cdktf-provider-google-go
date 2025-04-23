// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformtenantoauthidpconfig


type IdentityPlatformTenantOauthIdpConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/identity_platform_tenant_oauth_idp_config#create IdentityPlatformTenantOauthIdpConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/identity_platform_tenant_oauth_idp_config#delete IdentityPlatformTenantOauthIdpConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/identity_platform_tenant_oauth_idp_config#update IdentityPlatformTenantOauthIdpConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

