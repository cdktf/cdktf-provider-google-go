// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformtenantinboundsamlconfig


type IdentityPlatformTenantInboundSamlConfigSpConfig struct {
	// Callback URI where responses from IDP are handled. Must start with 'https://'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/identity_platform_tenant_inbound_saml_config#callback_uri IdentityPlatformTenantInboundSamlConfig#callback_uri}
	CallbackUri *string `field:"required" json:"callbackUri" yaml:"callbackUri"`
	// Unique identifier for all SAML entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.5.0/docs/resources/identity_platform_tenant_inbound_saml_config#sp_entity_id IdentityPlatformTenantInboundSamlConfig#sp_entity_id}
	SpEntityId *string `field:"required" json:"spEntityId" yaml:"spEntityId"`
}

