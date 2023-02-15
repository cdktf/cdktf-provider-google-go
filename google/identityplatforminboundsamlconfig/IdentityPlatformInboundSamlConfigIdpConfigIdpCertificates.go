package identityplatforminboundsamlconfig


type IdentityPlatformInboundSamlConfigIdpConfigIdpCertificates struct {
	// The IdP's x509 certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/identity_platform_inbound_saml_config#x509_certificate IdentityPlatformInboundSamlConfig#x509_certificate}
	X509Certificate *string `field:"optional" json:"x509Certificate" yaml:"x509Certificate"`
}

