// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatforminboundsamlconfig


type IdentityPlatformInboundSamlConfigIdpConfigIdpCertificates struct {
	// The IdP's x509 certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.24.0/docs/resources/identity_platform_inbound_saml_config#x509_certificate IdentityPlatformInboundSamlConfig#x509_certificate}
	X509Certificate *string `field:"optional" json:"x509Certificate" yaml:"x509Certificate"`
}

