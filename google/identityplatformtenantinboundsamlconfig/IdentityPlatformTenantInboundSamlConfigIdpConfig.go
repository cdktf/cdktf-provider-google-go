package identityplatformtenantinboundsamlconfig


type IdentityPlatformTenantInboundSamlConfigIdpConfig struct {
	// idp_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_tenant_inbound_saml_config#idp_certificates IdentityPlatformTenantInboundSamlConfig#idp_certificates}
	IdpCertificates interface{} `field:"required" json:"idpCertificates" yaml:"idpCertificates"`
	// Unique identifier for all SAML entities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_tenant_inbound_saml_config#idp_entity_id IdentityPlatformTenantInboundSamlConfig#idp_entity_id}
	IdpEntityId *string `field:"required" json:"idpEntityId" yaml:"idpEntityId"`
	// URL to send Authentication request to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_tenant_inbound_saml_config#sso_url IdentityPlatformTenantInboundSamlConfig#sso_url}
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Indicates if outbounding SAMLRequest should be signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/identity_platform_tenant_inbound_saml_config#sign_request IdentityPlatformTenantInboundSamlConfig#sign_request}
	SignRequest interface{} `field:"optional" json:"signRequest" yaml:"signRequest"`
}

