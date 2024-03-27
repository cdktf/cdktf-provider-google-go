// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatforminboundsamlconfig


type IdentityPlatformInboundSamlConfigSpConfig struct {
	// Callback URI where responses from IDP are handled. Must start with 'https://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/identity_platform_inbound_saml_config#callback_uri IdentityPlatformInboundSamlConfig#callback_uri}
	CallbackUri *string `field:"optional" json:"callbackUri" yaml:"callbackUri"`
	// Unique identifier for all SAML entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/identity_platform_inbound_saml_config#sp_entity_id IdentityPlatformInboundSamlConfig#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
}

