// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatforminboundsamlconfig


type IdentityPlatformInboundSamlConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/identity_platform_inbound_saml_config#create IdentityPlatformInboundSamlConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/identity_platform_inbound_saml_config#delete IdentityPlatformInboundSamlConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/identity_platform_inbound_saml_config#update IdentityPlatformInboundSamlConfig#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

